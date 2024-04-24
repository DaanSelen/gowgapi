package wgsqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	wgdb *sql.DB
)

func InitDatabase() bool {
	var err error
	wgdb, err = sql.Open("sqlite3", "wgdb.db")
	if err != nil {
		log.Fatal(err)
	}

	err = wgdb.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
		return false
	} else {
		log.Println("WGSqlite persistency succesfully connected.")
		return true
	}
}
