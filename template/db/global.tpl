package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/jtao539/autocode/template/config"
	"github.com/jtao539/sqlxp"

	_ "github.com/go-sql-driver/mysql"
)

var Pro *sqlxp.SqlxP

func Init() {
	d := config.Conf.DB
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", d.DBUser, d.DBPass, d.DBHost, d.DBPort, d.DBName)
	connect("mysql", conStr)
}

func connect(driveName, dialect string) {
	db, err := sqlx.Open(driveName, dialect)
	if err != nil {
		panic(fmt.Sprintf("Error in connect db:%s", err.Error()))
	}
	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("Error in connect db:%s", err.Error()))
	}
	Pro = &sqlxp.SqlxP{DB: db}
}

func Close() {
	Pro.DB.Close()
}
