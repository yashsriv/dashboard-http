// main.go
package main

import (
	"github.com/yashsriv/dashboard-http/router"

	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	"gopkg.in/kataras/iris.v5"
)

func main() {

	iris.Config.Gzip = true
	iris.Use(logger.New())
	iris.Use(recovery.New())

	router.DashboardRoute()

	iris.Listen(":8080")

}
