package router

import (
	"gopkg.in/kataras/iris.v5"
)

func DashboardRoute() {

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Text(iris.StatusAccepted, "")
	})

}
