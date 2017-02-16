package router

import (
	"gopkg.in/kataras/iris.v5"
)

// DashboardRoute - Function to set up Iris Router
func DashboardRoute() {

	iris.Get("/", func(ctx *iris.Context) {
		_ = ctx.Text(iris.StatusAccepted, "")
	})

}
