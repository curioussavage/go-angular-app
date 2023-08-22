package controllers

import (
	"net/http"

	"github.com/curioussavage/integra/models"
	"github.com/labstack/echo/v4"
)

// @Summary Get a list of users
// @Description get users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsersController(c echo.Context) error {
	// TODO pagination (probably skip on this for time)
	users := []models.User{
		{
			UserName:   "user1",
			FirstName:  "john",
			LastName:   "doe",
			Email:      "foo@b.com",
			UserStatus: models.Active,
		},
	}
	return c.JSON(http.StatusOK, users)
}

// @Summary Delete a user
// @Description Delete a user
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /user [delete]
func DeleteUsercontroller(c echo.Context) error {
	return c.String(http.StatusOK, "delete")
}

// @Summary Update a user
// @Description Update a user
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /user/:id [post]
func UpdateUsercontroller(c echo.Context) error {
	return c.String(http.StatusOK, "update")
}

// @Summary create a user
// @Description create a user
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /user [post]
func CreateUsercontroller(c echo.Context) error {
	return c.String(http.StatusOK, "create")
}
