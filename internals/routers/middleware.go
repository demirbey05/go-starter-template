package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware" // Update import statement to use version 4
)

func InitGeneralMiddleware(e *echo.Echo, loggerConfig middleware.LoggerConfig) {
	e.Use(middleware.LoggerWithConfig(loggerConfig)) // Use the correct version of MiddlewareFunc
	e.Use(middleware.Recover())
}
