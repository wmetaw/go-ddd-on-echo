package config

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DBCon *gorm.DB
)

// new db connection returns db,err
func NewDBConnection() (*gorm.DB, error) {

	db, err := gorm.Open("mysql", GetDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDSN() string {
	user := getEnv("DB_USER", "root")
	password := getEnv("DB_PASSWORD", "root")
	db_name := "dev"
	opt := "charset=utf8&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(mysql-server:3306)/%s?%s", user, password, db_name, opt)
}

func getEnv(name, def string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}
	return def
}
