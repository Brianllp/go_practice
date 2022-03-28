package server

import (
	"github.com/Brianllp/go_practice/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	router := echo.New()

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.GET("/", controllers.Hello)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.GET("/entries", controllers.GetEntries)

	return router
}
