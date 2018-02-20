package interfaces

import (
	"github.com/labstack/echo"
)

// 各ルーティングに対するハンドラを設定
func Routes(e *echo.Echo) {
	e.GET("/users", Users)
	e.GET("/users/:id", UsersGet)
	e.POST("/users/:id", UsersUpdate)
}
