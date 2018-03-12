package main

import (
	"github.com/labstack/echo"
	"github.com/wmetaw/go-ddd-on-echo/config"
	"github.com/wmetaw/go-ddd-on-echo/interfaces/handlers"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// go test github.com/wmetaw/go-ddd-on-echo/test/

func TestUsersHandler(t *testing.T) {

	e := echo.New()
	handlers.Routes(e)

	// DB Connection
	var err error
	config.MysqlCon, err = config.NewMysqlConnection()
	if err != nil {
		panic(err)
	}
	defer config.MysqlCon.Close()

	// Users
	req := httptest.NewRequest(echo.GET, "/users", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// log テストは終了しない
	t.Logf("Log : %v", rec.Body)

	// Update
	f := make(url.Values)
	f.Set("name", "Sato")
	req = httptest.NewRequest(echo.POST, "/users/1", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// log テストは終了しない
	t.Logf("Log : %v", rec.Body)

	/*
		t.Fatal系が呼び出し元のテストメソッドの実行を即座に終了させるのに対し、
		t.Error系はテストを失敗扱いにするものの、処理はそのまま継続させる。
	*/
	//t.Fatal("Fatal")
	//t.Error("Error") // 表示されない
}
