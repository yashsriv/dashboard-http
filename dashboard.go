// main.go
package main

import (
	"github.com/yashsriv/dashboard-http/router"

	"github.com/iris-contrib/middleware/logger"
	"gopkg.in/kataras/iris.v5"
)

func main() {

	iris.Config.Gzip = true
	iris.Use(logger.New())

	router.DashboardRoute()

	iris.Listen(":8080")

}
