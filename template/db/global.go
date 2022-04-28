package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/jtao539/autocode"

	_ "github.com/go-sql-driver/mysql"
)

var GDB *saleDB
var Pro *autocode.SqlPro

func Init(userName, password, host, port, name string) {
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userName, password, host, port, name)
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
	GDB = &saleDB{DB: db}
	Pro = &autocode.SqlPro{DB: db}
}

type saleDB struct {
	DB *sqlx.DB
}

func Close() {
	GDB.DB.Close()
}
