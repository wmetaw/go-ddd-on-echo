package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wmetaw/go-ddd-on-echo/interfaces"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// handler
	interfaces.Routes(e)

	// start server
	e.Logger.Fatal(e.Start(":1323"))
}
