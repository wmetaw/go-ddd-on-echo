package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

var (
	DBCon *gorm.DB
)

// new db connection returns db,err
func NewDBConnection() (*gorm.DB, error) {

	user := getEnv("DB_USER", "root")
	password := getEnv("DB_PASSWORD", "")
	db_name := "test"
	opt := "charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@/%s?%s", user, password, db_name, opt)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getEnv(name, def string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}
	return def
}
