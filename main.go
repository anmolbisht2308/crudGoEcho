package main

import (
	"myapp/config"
	"myapp/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Connect to MongoDB
	config.ConnectDB()

	// Initialize routes
	routers.InitRoutes(e)
	routers.InitTodoRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
