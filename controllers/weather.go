package controllers

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/parnurzeal/gorequest"
	"github.com/yashsriv/dashboard-http/config"
	"gopkg.in/kataras/iris.v5"
)

// WeatherController for weather
func WeatherController(ctx *iris.Context) {
	val, err := config.RedisConnection.Get("weather").Result()
	if err != nil && err != redis.Nil {
		SendInternalServer(err, ctx)
		return
	}
	if err == redis.Nil {
		getWeather(ctx)
	} else {
		var v interface{}
		err := json.Unmarshal([]byte(val), &v)
		if err != nil {
			SendInternalServer(err, ctx)
			return
		}
		ctx.JSON(iris.StatusOK, v)
	}
}

func getWeather(ctx *iris.Context) {

	request := gorequest.New().Get("http://api.openweathermap.org/data/2.5/weather").
		Param("q", "Kanpur,in").
		Param("APPID", config.WeatherApiKey)

	resp, body, errs := request.End()
	if errs != nil {
		SendInternalServer(errs[0], ctx)
	}
	if resp.StatusCode == 200 {
		var v interface{}
		err := json.Unmarshal([]byte(body), &v)
		if err != nil {
			SendInternalServer(err, ctx)
			return
		}
		err = config.RedisConnection.Set("weather", body, time.Minute*10).Err()
		if err != nil {
			SendInternalServer(err, ctx)
			return
		}
		ctx.JSON(iris.StatusOK, v)
	} else {
		ctx.Text(iris.StatusInternalServerError, "Could Not Fetch")
	}

}
