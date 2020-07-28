package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB = nil

func GetDB() *sql.DB {
	if db != nil {
		return db
	}

	dbc, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/summoner?charset=utf8&parseTime=true")

	db = dbc

	if err != nil {
		fmt.Println("初始化DB失败", err)
	}

	return db
}
