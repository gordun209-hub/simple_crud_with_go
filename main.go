package main

import (
	r"echoapp/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//----------
// Handlers
//----------

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/users", r.GetAllUsers)
	e.POST("/users", r.CreateUser)
	e.GET("/users/:id", r.GetUser)
	e.PUT("/users/:id", r.UpdateUser)
	e.DELETE("/users/:id", r.DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
