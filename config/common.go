package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
)

const (
	// Environment Value
	GOENV = "GOENV"

	MYSQL_DB       = "MYSQL_DB"
	MYSQL_HOST     = "MYSQL_HOST"
	MYSQL_PORT     = "MYSQL_PORT"
	MYSQL_USER     = "MYSQL_USER"
	MYSQL_PASSWORD = "MYSQL_PASSWORD"

	MEMCACHE_HOST = "MEMCACHE_HOST"
	MEMCACHE_PORT = "MEMCACHE_PORT"

	REDIS_HOST     = "REDIS_HOST"
	REDIS_PORT     = "REDIS_PORT"
	REDIS_PASSWORD = "REDIS_PASSWORD"
)

func LoadEnv() {

	// 環境変数の設定がない場合はローカル設定
	if os.Getenv(GOENV) == "" {
		os.Setenv(GOENV, "development")
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv(GOENV)))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getEnv(name, def string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}
	return def
}
