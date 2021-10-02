package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlConfig struct {
	user     string
	password string
	ip       string
	database string
}

type DbConnections interface{}

func initMySqlConn() (db *sql.DB) {

	var config MySqlConfig = MySqlConfig{
		user:     "USER",
		password: "PASSWORD",
		ip:       "IP",
		database: "DATABASE",
	}

	var connString string = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		config.user,
		config.password,
		config.ip,
		config.database,
	)

	db, err := sql.Open("mysql", connString)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10)

	return db

}
