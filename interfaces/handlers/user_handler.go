package handlers

import (
	"github.com/labstack/echo"
	"github.com/wmetaw/go-ddd-on-echo/application"
	"net/http"
	"strconv"
)

func Users(c echo.Context) error {

	users, err := application.Users()
	if err != nil {
		return JsonError(c, err, &APIError{
			Message: "ユーザーが見つかりませんでした",
			Status:  http.StatusInternalServerError,
			Code:    ErrCodeText(CodeClientHoge),
		})
	}

	data := map[string]interface{}{"user": users}
	return c.JSON(http.StatusOK, data)
}

func UsersGet(c echo.Context) error {

	pid := c.Param("id")
	id, e := strconv.Atoi(pid)
	if e != nil || id <= 0 {
		return JsonError(c, e, &APIError{
			Message: "パラメーターが不正です",
			Status:  http.StatusInternalServerError,
			Code:    ErrCodeText(CodeClientHoge),
		})
	}

	user, err := application.UsersGet(id)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{"user": user}
	return c.JSON(http.StatusOK, data)
}

func UsersUpdate(c echo.Context) error {

	pid := c.Param("id")
	name := c.FormValue("name")
	id, e := strconv.Atoi(pid)
	if e != nil || id <= 0 {
		return JsonError(c, e, &APIError{
			Message: "パラメーターが不正です",
			Status:  http.StatusInternalServerError,
			Code:    ErrCodeText(CodeClientHoge),
		})
	}

	user, err := application.UsersUpdate(id, name)
	if err != nil {
		return JsonError(c, e, &APIError{
			Message: "更新に失敗しました",
			Status:  http.StatusInternalServerError,
			Code:    ErrCodeText(CodeClientHoge),
		})
	}

	data := map[string]interface{}{"user": user}
	return c.JSON(http.StatusOK, data)
}
