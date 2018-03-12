package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wmetaw/go-ddd-on-echo/config"
	"github.com/wmetaw/go-ddd-on-echo/interfaces/handlers"
)

func main() {

	// 環境変数をロード
	config.LoadEnv()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	// DB Connection
	var err error
	config.MysqlCon, err = config.NewMysqlConnection()
	if err != nil {
		panic(err)
	}
	defer config.MysqlCon.Close()

	// Memcached Connection
	config.MemcacheCon = config.NewMemcacheConnection()

	// Redis Connection
	config.RedisCon, err = config.NewRedisConnection()
	if err != nil {
		panic(err)
	}

	// migrate
	//library.Migrate()

	// handler
	handlers.Routes(e)

	// start server
	e.Logger.Fatal(e.Start(":1323"))
}
