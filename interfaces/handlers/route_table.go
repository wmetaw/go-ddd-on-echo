package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

// 各ルーティングに対するハンドラを設定
func Routes(e *echo.Echo) {
	e.GET("/users", Users)
	e.GET("/users/:id", UsersGet)
	e.POST("/users/:id", UsersUpdate)
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var (
		status = http.StatusInternalServerError
		msg    interface{}
	)

	if he, ok := err.(*echo.HTTPError); ok {
		status = he.Code
		msg = he.Message
		if he.Inner != nil {
			msg = fmt.Sprintf("%v, %v", err, he.Inner)
		}
	} else if c.Echo().Debug {
		msg = err.Error()
	} else {
		msg = http.StatusText(status)
	}
	if _, ok := msg.(string); ok {
		msg = map[string]interface{}{"status": status, "message": msg, "code": "1234"}
	}

	c.Echo().Logger.Error(err)

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == "HEAD" {
			err = c.NoContent(status)
		} else if c.Request().Header.Get("Content-Type") == "text/html" {
			// HTML
			err = c.File(fmt.Sprintf("%d.html", status))
		} else {
			err = c.JSON(status, msg)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}
