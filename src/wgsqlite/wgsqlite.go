package wgsqlite

// Package for SQLite data persistence and config storage.

import (
	"database/sql"
	"log"
	"strings"

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

func QueryAllInterfaces() []InterfaceQueryStruct {
	rows, _ := wgdb.Query("SELECT name, addr, port, out_interface, privkey FROM iface")
	var result []InterfaceQueryStruct

	defer rows.Close()
	for rows.Next() {
		var singleInterface InterfaceQueryStruct
		rows.Scan(&singleInterface.Name, &singleInterface.Address, &singleInterface.Port, &singleInterface.Out_Interface, &singleInterface.PrivKey)
		result = append(result, singleInterface)
	}

	return result
}

func QuieryPrivKey(ifaceName string) string {
	row := wgdb.QueryRow("SELECT privkey FROM iface WHERE name == ?;", ifaceName)
	var privKey string
	row.Scan(&privKey)
	return privKey
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

func checkDuplicateInterface(iface string) bool {
	row := wgdb.QueryRow("SELECT COUNT(*) FROM iface WHERE name == ?", iface)
	var rowCount int
	row.Scan(&rowCount)

	return rowCount > 0
}

func checkDuplicateNetwork(address, port string) bool {
	octets := strings.Split(address, ".")
	row := wgdb.QueryRow("SELECT COUNT(*) FROM iface WHERE addr LIKE @pattern", sql.Named("pattern", "%.%."+octets[2]+".%")) // Check the third octec for /24 subnets.
	var netRowCount int
	row.Scan(&netRowCount)

	row = wgdb.QueryRow("SELECT COUNT(*) FROM iface WHERE port == ?", port)
	var portRowCount int
	row.Scan(&portRowCount)

	return (netRowCount > 0 || portRowCount > 0)
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
