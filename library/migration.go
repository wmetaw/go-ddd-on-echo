package library

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
	"github.com/wmetaw/go-ddd-on-echo/config"
)

func Migrate() {
	db, _ := sql.Open("mysql", config.GetMysqlDSN())
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://config/migrate",
		"mysql",
		driver,
	)

	m.Steps(2)
}
