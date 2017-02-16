// main.go
package main

import (
	"fmt"

	"github.com/yashsriv/dashboard-http/router"

	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	"github.com/olebedev/config"
	"gopkg.in/kataras/iris.v5"
)

func main() {

	iris.Config.Gzip = true
	iris.Use(logger.New())
	iris.Use(recovery.New())

	router.DashboardRoute()

	cfg, _ := config.ParseYamlFile("./config.yml")
	cfg.EnvPrefix("DASHBOARD")

	// Can be set using DASHBOARD_HTTP_PORT environment variable
	port, _ := cfg.Int("http.port")

	iris.Listen(fmt.Sprintf(":%d", port))

}
