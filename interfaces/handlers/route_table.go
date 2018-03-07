package handlers

import (
	"bufio"
	"github.com/labstack/echo"
	"net"
	"net/http"
	"strconv"
	"strings"
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
	m, _ := Mem("memcached-server:11211")
	m.set("foo", []byte("aaaaa"))
	return c.JSON(http.StatusOK, map[string]interface{}{"data": pid})
}

func MemcacheGet(c echo.Context) error {
	m, _ := Mem("memcached-server:11211")
	res, _ := m.get("foo")
	return c.JSON(http.StatusOK, map[string]interface{}{"data": string(res)})
}

type Memcache struct {
	conn     net.Conn
	buffered bufio.ReadWriter
}

func Mem(addr string) (conn *Memcache, err error) {
	nc, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Memcache{
		conn:     nc,
		buffered: bufio.ReadWriter{Reader: bufio.NewReader(nc), Writer: bufio.NewWriter(nc)},
	}, err
}

func (mc *Memcache) get(key string) (result []byte, err error) {
	_, err = mc.buffered.WriteString("get " + key + "\n")
	if err == nil {
		err = mc.buffered.Flush()
		if err == nil {
			for {
				b, _, err := mc.buffered.ReadLine()
				l := string(b)
				if err == nil {
					if strings.HasPrefix(l, "END") {
						break
					}
					if strings.Contains(l, "ERROR") {
						panic("ERROR")
					}
					if !strings.HasPrefix(l, "VALUE") {
						result = append(result, l...)
						result = append(result, '\n')
					}
				} else {
					panic(err)
				}
			}
		} else {
			panic(err)
		}
	}
	return result, err
}

func (mc *Memcache) set(key string, value []byte) (err error) {
	_, err = mc.buffered.WriteString("set " + key + " 0 0 " + strconv.Itoa(len(value)) + "\r\n")
	if err == nil {
		v := append(value, "\r\n"...)
		_, err = mc.buffered.Write(v)
		if err != nil {
			panic(err)
		}
		err = mc.buffered.Flush()
		if err == nil {
			mc.buffered.ReadLine()

		}
	}
	return err
}
