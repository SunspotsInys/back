package db

import (
	"fmt"

	"github.com/SunspotsInys/thedoor/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
	// logger = logs.Logger
)

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		configs.Conf.DBName, configs.Conf.DBPasswd, configs.Conf.DBHost,
		configs.Conf.DBPort, configs.Conf.DBName,
	)
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(3)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
