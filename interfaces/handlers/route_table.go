package handlers

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis"
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

	// redis test
	e.GET("/redis/:id", RedisGet)
	e.POST("/redis/:id", RedisSet)
}

func RedisSet(c echo.Context) error {

	pid := c.Param("id")

	if err := config.RedisCon.Set(pid, "fugavalue", 0).Err(); err != nil {
		return JsonError(c, err, &APIError{
			Message: "Redis Can not Set : " + pid,
			Status:  http.StatusInternalServerError,
			Code:    ErrCodeText(CodeClientHoge),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": pid})
}

func RedisGet(c echo.Context) error {

	pid := c.Param("id")

	// zadd FIFA 6 holland 6 colombia 2 japan 4 taiwan 3 china 1 germany 2 argentina
	cmd := config.RedisCon.ZScore("FIFA", "japan")
	cmdint := config.RedisCon.ZCount("FIFA", fmt.Sprint(cmd.Val()+1), "+inf")

	val, err := config.RedisCon.Get(pid).Result()
	if err == redis.Nil {
		// fmt.Println("hogekey does not exist")
	} else if err != nil {
		return JsonError(c, err, &APIError{
			Message: "Redis Can not Get : " + pid,
			Status:  http.StatusInternalServerError,
			Code:    ErrCodeText(CodeClientHoge),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": cmdint.Val() + 1, "data2": val})
}

func MemcacheSet(c echo.Context) error {

	pid := c.Param("id")

	if err := config.MemcacheCon.Set(&memcache.Item{Key: "foo", Value: []byte(pid)}); err != nil {
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

	it, err := config.MemcacheCon.Get("foo")
	if err != nil {
		return JsonError(c, err, &APIError{
			Message: "Can not Get : " + pid,
			Status:  http.StatusInternalServerError,
			Code:    ErrCodeText(CodeClientHoge),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": string(it.Value)})
}
