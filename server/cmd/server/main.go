package main

import (
	"github.com/curioussavage/integra/controllers"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/curioussavage/integra/docs"
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

	e.GET("/api/v1/users", controllers.GetUsersController)
	e.POST("/api/v1/user", controllers.CreateUsercontroller)
	e.POST("/api/v1/user/:id", controllers.UpdateUsercontroller)
	e.DELETE("/api/v1/user/:id", controllers.DeleteUsercontroller)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
