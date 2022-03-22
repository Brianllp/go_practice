package main

import (
	"github.com/Brianllp/go_practice/controllers"
	"github.com/Brianllp/go_practice/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	database.ConnectDB()
	defer database.CloseDB()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", controllers.Hello)
	e.GET("/users", controllers.GetUsers)
	e.GET("/users/:id", controllers.GetUser)
	e.GET("/entries", controllers.GetEntries)

	e.Logger.Fatal(e.Start(":3030"))
}
