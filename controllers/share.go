package controllers

import (
	"github.com/yashsriv/dashboard-http/config"
	"gopkg.in/kataras/iris.v5"
)

func GetPosts(ctx *iris.Context) {
	response, err := config.Share.Call("get")
	if err != nil {
		SendInternalServer(err, ctx)
	}
	var posts interface{}
	response.GetObject(&posts)

	_ = ctx.JSON(iris.StatusOK, posts)
}
