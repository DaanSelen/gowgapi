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

func InitDatabase() {
	var err error
	wgdb, err = sql.Open("sqlite3", "wgdb.db")
	if err != nil {
		log.Fatal("Failed to create database", err)
	}

	err = wgdb.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	} else {
		log.Println("WGSqlite persistency succesfully connected.")
		setupTables()
	}
}

func QueryUser(username string) UserQueryStruct {
	row := wgdb.QueryRow("SELECT username, password, role, salt FROM account WHERE username == ?;", username)
	var result UserQueryStruct
	row.Scan(&result.Username, &result.Password, &result.Role, &result.Salt)
	return result
}

func CheckEmptyAccountTable() bool {
	row := wgdb.QueryRow("SELECT COUNT(*) FROM account")
	var rowCount int
	row.Scan(&rowCount)

	return rowCount == 0
}

func checkDuplicateUser(username string) bool {
	row := wgdb.QueryRow("SELECT COUNT(*) FROM account WHERE username == ?", username)
	var rowCount int
	row.Scan(&rowCount)

	return rowCount > 0
}

func setupTables() {
	var count int = 0
	for x := range tableQueries { // Create all tables defined in the function call.
		_, err := wgdb.Exec(tableQueries[x])
		if err != nil {
			log.Fatal("Failed getting database ready:", err)
		} else {
			count++
		}
	}
	if count == len(tableQueries) {
		log.Println("Tables successfully inserted.")
	}
}
