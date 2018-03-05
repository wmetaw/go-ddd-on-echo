package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wmetaw/go-ddd-on-echo/config"
	"github.com/wmetaw/go-ddd-on-echo/interfaces/handlers"
)

func main() {

	// migrate
	//library.Migrate()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	// DB Connection
	var err error
	config.DBCon, err = config.NewDBConnection()
	if err != nil {
		panic(err)
	}
	defer config.DBCon.Close()

	// handler
	handlers.Routes(e)

	// start server
	e.Logger.Fatal(e.Start(":1323"))
}
