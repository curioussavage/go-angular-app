package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	return c.String(http.StatusOK, "get")
}

func DeleteUsercontroller(c echo.Context) error {
	return c.String(http.StatusOK, "delete")
}

func UpdateUsercontroller(c echo.Context) error {
	return c.String(http.StatusOK, "update")
}

func CreateUsercontroller(c echo.Context) error {
	return c.String(http.StatusOK, "create")
}
