package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	MysqlCon *gorm.DB
)

// new db connection returns db,err
func NewMysqlConnection() (*gorm.DB, error) {

	db, err := gorm.Open("mysql", GetMysqlDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetMysqlDSN() string {

	// 環境変数から値を取得。なければデフォルト値をセット
	user := getEnv(MYSQL_USER, "root")
	pwd := getEnv(MYSQL_PASSWORD, "root")
	host := getEnv(MYSQL_HOST, "mysql-server")
	port := getEnv(MYSQL_PORT, "3306")
	db := getEnv(MYSQL_DB, "dev")
	opt := "charset=utf8&parseTime=True&loc=Asia%2FTokyo"

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, pwd, host, port, db, opt)
}
