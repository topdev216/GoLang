package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func SetSqlite3(conn string) {
	db = sqlx.MustConnect("sqlite3", "./database.sqlite3")
}

func GetSqlite3() *sqlx.DB {
	return db
}
