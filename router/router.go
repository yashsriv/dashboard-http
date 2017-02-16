package router

import (
	"gopkg.in/kataras/iris.v5"

	"github.com/yashsriv/dashboard-http/controllers"
)

// DashboardRoute - Function to set up Iris Router
func DashboardRoute() {

	iris.Get("/", func(ctx *iris.Context) {
		_ = ctx.Text(iris.StatusAccepted, "")
	})

	iris.Get("/user/me", controllers.IsAuthenticated, func(ctx *iris.Context) {
		_ = ctx.JSON(iris.StatusOK, iris.Map{"username": ctx.GetCookie("username")})
	})

	iris.Post("/user/login", controllers.Login)

}
