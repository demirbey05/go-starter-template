package routers

import (
	"github.com/demirbey05/golang-starter/internals/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes() *echo.Echo {
	e := echo.New()
	InitGeneralMiddleware(e, middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, ip=${remote_ip}, uri=${uri}, status=${status}\n",
	})
	e.GET("/", controllers.HelloHandler)
	return e
}
