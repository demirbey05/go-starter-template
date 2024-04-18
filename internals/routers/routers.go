package routers

import (
	"github.com/demirbey05/golang-starter/internals/controllers"
	"github.com/labstack/echo/v4"
)

func InitRoutes() *echo.Echo {
	e := echo.New()
	e.GET("/", controllers.HelloHandler)
	return e
}
