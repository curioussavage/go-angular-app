package controllers

import (
	"log"
	"net/http"
	"strconv"

	apperrors "github.com/curioussavage/integra/errors"
	"github.com/curioussavage/integra/models"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	GetUsers(query models.UserFilters) ([]models.User, error)
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
// @Param id query int false "A user id to filter by"
// @Success 200 {array} models.User
// @Router /users [get]
func (ctl *UserController) GetUsersController(c echo.Context) error {
	var userFilters models.UserFilters
	err := c.Bind(&userFilters)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	// TODO pagination (probably skip on this for time)
	users, err := ctl.userService.GetUsers(userFilters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, struct{}{})
	}
	return c.JSON(http.StatusOK, users)
}

// @Summary Delete a user
// @Description Delete a user
// @Accept json
// @Produce json
// @Param id path int true "id of user to delete"
// @Success 200 {string} string "ok"
// @Router /user/{id} [delete]
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
	return c.JSON(http.StatusOK, "ok")
}

// @Summary Update a user
// @Description Update a user
// @Accept json
// @Produce json
// @Param id path int true "id of user to update"
// @Param userUpdate body models.UserUpdateForm true "user data to update"
// @Success 200 {object} models.User
// @Failure 400 {array} apperrors.ValidationErrors
// @Router /user/{id} [patch]
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

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
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
// @Failure 400 {array} apperrors.ValidationErrors
// @Router /user [post]
func (ctl *UserController) CreateUsercontroller(c echo.Context) error {
	var user models.UserCreationForm
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad request")
	}

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newUser, err := ctl.userService.CreateUser(user)
	if err != nil {
		if _, ok := err.(*models.UsernameTakenError); ok {
			err := apperrors.ValidationErrors{apperrors.ValidationError{
				Field:   "userName",
				Message: "username not available.",
			}}
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, newUser)
}
