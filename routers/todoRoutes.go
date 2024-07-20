package routers

import (
	"myapp/controllers"

	"github.com/labstack/echo/v4"
)

func InitTodoRoutes(e *echo.Echo) {
	e.POST("/todos", controllers.CreateTodo)
	e.GET("/todos", controllers.GetTodos)
	e.GET("/todos/:id", controllers.GetTodo)
	e.PUT("/todos/:id", controllers.UpdateTodo)
	e.DELETE("/todos/:id", controllers.DeleteTodo)
}
