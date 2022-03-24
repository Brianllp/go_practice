package controllers

import (
	"net/http"

	"github.com/Brianllp/go_practice/models"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users := models.GetUsers()
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	user := models.FindUserByID(id)
	return c.JSON(http.StatusOK, user)
}
