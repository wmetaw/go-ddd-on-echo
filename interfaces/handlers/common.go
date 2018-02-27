package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

var errorCode = map[int]string{}

type APIError struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Code    string      `json:"code"`
}

func JsonError(c echo.Context, err error, a *APIError) error {

	// 通信自体は完了しているので200
	c.JSON(http.StatusOK, a)
	return err
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var (
		status = http.StatusInternalServerError
		msg    interface{}
		code   = CodeServerHoge
	)

	if he, ok := err.(*echo.HTTPError); ok {
		status = he.Code
		msg = http.StatusText(status)

		// middleware error handling
		if he.Inner != nil {
			msg = fmt.Sprintf("%v, %v", err, he.Inner)
		}
	} else if c.Echo().Debug {
		msg = err.Error()
	} else {
		msg = http.StatusText(status)
	}

	if _, ok := msg.(string); ok {
		msg = map[string]interface{}{"status": status, "message": msg, "code": code}
	}
	c.Echo().Logger.Error(err)

	// Send response
	if !c.Response().Committed {
		c.Response().Header().Set("Content-Status", "333")
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
