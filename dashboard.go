// main.go
package main

import (
	"fmt"

	"github.com/yashsriv/dashboard-http/router"

	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	"github.com/olebedev/config"
	"gopkg.in/kataras/iris.v5"
)

func main() {

	iris.Config.Gzip = true
	iris.Use(logger.New())
	iris.Use(recovery.New())
	iris.Use(cors.Default())

	// log http errors
	iris.OnError(iris.StatusNotFound, myCorsMiddleware)

	router.DashboardRoute()

	cfg, err := config.ParseYamlFile("./config.yml")
	if err != nil {
		panic(err)
	}
	cfg.EnvPrefix("DASHBOARD")

	// Can be set using DASHBOARD_HTTP_PORT environment variable
	port, err := cfg.Int("http.port")
	if err != nil {
		panic(err)
	}

	iris.Listen(fmt.Sprintf(":%d", port))

}

// myCorsMiddleware for handling OPTIONS requests
func myCorsMiddleware(ctx *iris.Context) {

	if ctx.MethodString() == "OPTIONS" {
		ctx.SetHeader("Access-Control-Allow-Origin", "*")
		ctx.SetHeader("Access-Control-Allow-Headers", "content-type")
		err := ctx.Text(200, "")
		if err != nil {
			panic(err)
		}
	} else {
		errorLogger := logger.New()
		errorLogger.Serve(ctx)
	}

}
