package handlers

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wmetaw/go-ddd-on-echo/config"
	"net/http"
	"strconv"
	"time"
)

// 各ルーティングに対するハンドラを設定
func Routes(ex *echo.Echo) {

	// Login
	ex.POST("/login", login)

	// 以下、JWT認証が必要なエンドポイント
	e := ex.Group("")

	// 署名アルゴリズムの指定
	// 署名アルゴリズムが HMAC SHA なので鍵の文字列を指定
	// RSA や ECDSA など公開鍵暗号の場合は秘密鍵と公開鍵のファイルのパスを指定する
	e.Use(middleware.JWT(config.GetJWTKey()))
	e.Use(validTokenByIssueAt)

	e.GET("/restricted", restricted)

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

func login(c echo.Context) error {

	user := c.FormValue("user")
	pass := c.FormValue("pass")

	var err error

	if user == "admin" && pass == "pass" {

		// Create Token(HMAC SHA)
		token := jwt.New(jwt.SigningMethodHS256)

		// Set Claim
		claims := token.Claims.(jwt.MapClaims)
		claims["user"] = user
		claims["iat"] = time.Now().Unix()
		claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

		// tokenを秘密鍵でエンコード
		t, err := token.SignedString(config.GetJWTKey())
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{"token": t})
	}

	return JsonError(c, err, &APIError{
		Message: "id or password cannot be recognized.",
		Status:  http.StatusUnauthorized,
		Code:    ErrCodeText(CodeClientHoge),
	})
}

// Token force reset by IssueAt
func validTokenByIssueAt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// config.RedisCon.Set("iat", time.Now().Unix(), 0).Err()
		val, err := config.RedisCon.Get("iat").Result()

		// cache miss (値がなければ次のmiddlewareへ)
		if err == redis.Nil {
			return next(c)
		}
		redis_iat, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return JsonError(c, err, &APIError{
				Message: "Can not parse iat",
				Status:  http.StatusInternalServerError,
				Code:    ErrCodeText(CodeClientHoge),
			})
		}

		// Tokenを復号し値を取り出す
		// contextにセットされたclaim取得。storeのキーはデフォルトでuserになっている(JWTミドルウェアセット時にkey名を変更可)
		// 取得した値をToken型にキャスト
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		// interface型をキャスト
		if iat, ok := claims["iat"].(int64); ok {
			if redis_iat > iat {
				return JsonError(c, err, &APIError{
					Message: "Token Force Reset",
					Status:  http.StatusUnauthorized,
					Code:    ErrCodeText(CodeClientHoge),
				})
			}
		}

		return next(c)
	}
}

func restricted(c echo.Context) error {

	// Tokenを複合し値を取り出す
	// contextにセットされたclaim取得。storeのキーはデフォルトでuserになっている
	// JWTミドルウェアセット時にkey名を変更可
	// 取得した値をToken型にキャスト
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	u := claims["user"].(string)
	return c.String(http.StatusOK, "welcome! "+u)
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
