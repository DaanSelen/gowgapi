package wgsqlite

// Package for SQLite data persistence and config storage.

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
		setupTables(userTab, ifaceTab, clientTab)
		return true
	}
}

func QueryUser(username string) UserQueryStruct {
	log.Printf("SELECT * FROM user WHERE username == '%s';", username)
	row := wgdb.QueryRow("SELECT username, password, role, salt FROM user WHERE username == ?;", username)
	var result UserQueryStruct
	row.Scan(&result.Username, &result.Password, &result.Role, &result.Salt)
	return result
}

func setupTables(tables ...string) {
	var count int = 0
	for x := range tables { // Create all tables defined in the function call.
		_, err := wgdb.Exec(tables[x])
		if err != nil {
			log.Println(err)
		} else {
			count++
		}
	}
	if count == len(tables) {
		log.Println("Tables successfully inserted.")
	}
}
