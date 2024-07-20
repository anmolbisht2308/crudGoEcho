package routers

import (
	"myapp/controllers"

	"github.com/labstack/echo/v4"
)

// InitRoutes initializes the routes
func InitRoutes(e *echo.Echo) {
	e.GET("/", controllers.HelloWorld)
}
