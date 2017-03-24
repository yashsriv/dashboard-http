package controllers

import (
	"encoding/json"

	"github.com/yashsriv/dashboard-http/models"

	"github.com/parnurzeal/gorequest"
	"github.com/yashsriv/dashboard-http/config"
	"gopkg.in/kataras/iris.v5"
)

type fbJSON struct {
	AccessToken string `json:"access_token" xml:"access_token" form:"access_token"`
}

type fbResponse struct {
	AppID       string `json:"app_id"`
	Application string `json:"application"`
	IsValid     bool   `json:"is_valid"`
	UserID      string `json:"user_id"`
}

type fbResponseData struct {
	Data fbResponse `json:"data"`
}

// AddFacebook adds user's facebook access token
func AddFacebook(ctx *iris.Context) {
	username := ctx.RequestHeader("X-Username-Header")
	fbInfo := fbJSON{}
	err := ctx.ReadJSON(&fbInfo)
	if err != nil {
		_ = ctx.Text(iris.StatusBadRequest, err.Error())
	} else {
		if verify(fbInfo.AccessToken) {
			// Do something
			var users []models.User
			err := config.DatabaseConnection.Q().
				Where("username = ?", username).
				All(&users)
			if err != nil {
				SendInternalServer(err, ctx)
				return
			}
			if len(users) > 0 {
				user := users[0]
				user.Fbtoken = fbInfo.AccessToken
				config.DatabaseConnection.Update(&user)
				_ = ctx.Text(iris.StatusOK, "")
			} else {
				_ = ctx.Text(iris.StatusNotFound, "")
			}

		} else {
			_ = ctx.Text(iris.StatusNotAcceptable, "")
		}
	}
}

func verify(token string) bool {
	request := gorequest.New().Get("https://graph.facebook.com/debug_token").
		Param("input_token", token).
		Param("access_token", config.FacebookAccessToken)

	resp, body, errs := request.End()
	if errs != nil {
		return false
	}
	if resp.StatusCode == 200 {
		var r fbResponseData
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			return false
		}
		if r.Data.IsValid {
			return true
		}
	}
	return false
}
