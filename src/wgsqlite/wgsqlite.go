package wgsqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	wgdb *sql.DB
)

func InitDatabase() {
	var err error
	wgdb, err = sql.Open("sqlite3", "wgdb.db")
	if err != nil {
		log.Fatal(err)
	}

	err = wgdb.Ping() //Test the DB Connection.
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	} else {
		log.Println("NMTAS SQLite3 Database, Ready for connections.")
	}
}
