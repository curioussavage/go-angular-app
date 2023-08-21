package main

import (
	"github.com/curioussavage/integra/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// TODO take in config
	// TODO graceful shutdown

	e.GET("/api/v1/users", controllers.GetUsersController)
	e.POST("/api/v1/user", controllers.CreateUsercontroller)
	e.POST("/api/v1/user/:id", controllers.UpdateUsercontroller)
	e.DELETE("/api/v1/user/:id", controllers.DeleteUsercontroller)

	e.Logger.Fatal(e.Start(":1323"))
}
