package main

import (
	"database/sql"
	"log"

	"github.com/curioussavage/integra/controllers"
	"github.com/curioussavage/integra/models"
	"github.com/curioussavage/integra/validation"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/curioussavage/integra/docs"
	_ "github.com/mattn/go-sqlite3"
)

// @title User API
// @version 1.0
// @description This is an API for managing users.
//
// @contact.name spiderman
// @contact.url http://super.spiderman
// @contact.email support@spiderman.com
//
// @BasePath /api/v1
func main() {
	e := echo.New()
	// TODO take in config
	// TODO graceful shutdown

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}
	userService := models.UserService{DB: db}
	userController := controllers.NewUserController(&userService)

	e.Validator = &validation.CustomValidator{Validator: validator.New()}

	// TODO this should only be done in dev
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/api/v1/users", userController.GetUsers)
	e.POST("/api/v1/user", userController.CreateUser)
	e.PATCH("/api/v1/user/:id", userController.UpdateUser)
	e.DELETE("/api/v1/user/:id", userController.DeleteUser)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
