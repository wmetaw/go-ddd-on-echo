package interfaces

import (
	"github.com/labstack/echo"
	"github.com/wmetaw/go-ddd-on-echo/application"
	"net/http"
	"strconv"
)

func Users(c echo.Context) error {

	users, err := application.Users()
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{"user": users}
	return c.JSON(http.StatusOK, data)
}

func UsersGet(c echo.Context) error {

	pid := c.Param("id")
	if pid == "" {
		return c.JSON(http.StatusInternalServerError, "no value for param requested")
	}
	id, _ := strconv.Atoi(pid)

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
	if pid == "" {
		return c.JSON(http.StatusInternalServerError, "no value for param requested")
	}
	id, _ := strconv.Atoi(pid)

	user, err := application.UsersUpdate(id, name)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{"user": user}
	return c.JSON(http.StatusOK, data)
}
