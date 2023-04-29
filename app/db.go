package app

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

var DB *sql.DB

func InitDB() {
	driver := revel.Config.StringDefault("db.driver", "mysql")
	host := revel.Config.StringDefault("db.host", "")
	port := revel.Config.StringDefault("db.port", "3306")
	name := revel.Config.StringDefault("db.name", "")
	user := revel.Config.StringDefault("db.user", "")
	password := revel.Config.StringDefault("db.password", "")

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, name)

	db, err := sql.Open(driver, connString)
	if err != nil {
		revel.AppLog.Errorf("DB connection failed: %s", err.Error())
		return
	}

	err = db.Ping()
	if err != nil {
		revel.AppLog.Errorf("DB connection failed: %s", err.Error())
		return
	}

	DB = db
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
