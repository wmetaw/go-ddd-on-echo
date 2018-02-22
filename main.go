package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wmetaw/go-ddd-on-echo/interfaces/handlers"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	// handler
	handlers.Routes(e)

	// start server
	e.Logger.Fatal(e.Start(":1323"))
}
