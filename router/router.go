package router

import (
	"gopkg.in/kataras/iris.v5"

	"github.com/yashsriv/dashboard-http/controllers"
)

// DashboardRoute - Function to set up Iris Router
func DashboardRoute() {

	user := iris.Party("/user")
	user.Get("/me", controllers.IsAuthenticated, controllers.CurrentUser)
	user.Post("/login", controllers.Login)
	user.Post("/facebook", controllers.IsAuthenticated, controllers.AddFacebook)

}
