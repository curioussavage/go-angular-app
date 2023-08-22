package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/curioussavage/integra/models"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	GetUsers() ([]models.User, error)
	CreateUser(user models.UserCreationForm) (models.User, error)
	DeleteUser(userID int) error
	UpdateUser(userID int, user models.UserUpdateForm) (models.User, error)
}

type UserController struct {
	userService UserService
}

func NewUserController(userService UserService) UserController {
	return UserController{userService: userService}
}

// @Summary Get a list of users
// @Description get users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func (ctl *UserController) GetUsersController(c echo.Context) error {
	// TODO pagination (probably skip on this for time)
	users, err := ctl.userService.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, struct{}{})
	}
	return c.JSON(http.StatusOK, users)
}

// @Summary Delete a user
// @Description Delete a user
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /user/:id [delete]
func (ctl *UserController) DeleteUsercontroller(c echo.Context) error {
	userIdString := c.Param("id")
	userID, err := strconv.Atoi(userIdString)
	if err != nil {
		// TODO add an err response type
		return c.String(http.StatusBadRequest, "Bad request: id must be an integer")
	}

	if err := ctl.userService.DeleteUser(userID); err != nil {
		return c.String(http.StatusInternalServerError, "Encountered an error while deleting user")

	}
	return c.String(http.StatusOK, "ok")
}

// @Summary Update a user
// @Description Update a user
// @Accept json
// @Produce json
// @Param user body models.UserUpdateForm true "user data to update"
// @Success 200 {object} models.User
// @Router /user/:id [patch]
func (ctl *UserController) UpdateUsercontroller(c echo.Context) error {
	var user models.UserUpdateForm
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	userIdString := c.Param("id")
	userID, err := strconv.Atoi(userIdString)
	if err != nil {
		// TODO add an err response type
		return c.String(http.StatusBadRequest, "Bad request: id must be an integer")
	}

	updatedUser, err := ctl.userService.UpdateUser(userID, user)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Could not update user")
	}
	return c.JSON(http.StatusOK, updatedUser)
}

// @Summary create a user
// @Description create a user
// @Accept json
// @Produce json
// @Param user body models.UserCreationForm true "user to create"
// @Success 200 {object} models.User
// @Router /user [post]
func (ctl *UserController) CreateUsercontroller(c echo.Context) error {
	var user models.UserCreationForm
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newUser, err := ctl.userService.CreateUser(user)
	if err != nil {
		if _, ok := err.(*models.UsernameTakenError); ok {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, newUser)
}
