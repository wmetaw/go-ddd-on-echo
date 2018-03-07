package handlers

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/labstack/echo"
	"github.com/wmetaw/go-ddd-on-echo/config"
	"net/http"
)

// 各ルーティングに対するハンドラを設定
func Routes(e *echo.Echo) {
	e.GET("/users", Users)
	e.GET("/users/:id", UsersGet)
	e.POST("/users/:id", UsersUpdate)

	// memcached test
	e.GET("/memcache/:id", MemcacheGet)
	e.POST("/memcache/:id", MemcacheSet)
}

func MemcacheSet(c echo.Context) error {

	pid := c.Param("id")

	if err := config.MCCon.Set(&memcache.Item{Key: "foo", Value: []byte(pid)}); err != nil {
		return JsonError(c, err, &APIError{
			Message: "Can not Set : " + pid,
			Status:  http.StatusInternalServerError,
			Code:    ErrCodeText(CodeClientHoge),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": pid})
}

func MemcacheGet(c echo.Context) error {

	pid := c.Param("id")

	it, err := config.MCCon.Get("foo")
	if err != nil {
		return JsonError(c, err, &APIError{
			Message: "Can not Get : " + pid,
			Status:  http.StatusInternalServerError,
			Code:    ErrCodeText(CodeClientHoge),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": string(it.Value)})
}
