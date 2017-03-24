// main.go
package main

import (
	"fmt"

	"github.com/yashsriv/dashboard-http/config"
	"github.com/yashsriv/dashboard-http/router"

	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	"gopkg.in/kataras/iris.v5"
)

func main() {

	iris.Config.Gzip = true
	iris.Config.LoggerPreffix = "[dashboard-http] "

	iris.Use(logger.New())
	iris.Use(recovery.New())

	// log http errors
	iris.OnError(iris.StatusNotFound, myCorsMiddleware)

	config.InitConfig()

	router.DashboardRoute()

	iris.Listen(fmt.Sprintf(":%d", config.HTTPPort))

}

// myCorsMiddleware for handling OPTIONS requests
func myCorsMiddleware(ctx *iris.Context) {

	errorLogger := logger.New()
	errorLogger.Serve(ctx)
	_ = ctx.Text(iris.StatusNotFound, "")

}
