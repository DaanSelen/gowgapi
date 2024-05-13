package wgsqlite

import (
	"gowgapi/wgcrypt"
	"log"
)

func SaveAccount(username, password, role string) {
	prep, err := wgdb.Prepare("INSERT INTO account (username, password, role, salt) VALUES (?, ?, ?, ?);")
	if err != nil {
		log.Println(err)
	}
	defer prep.Close()

	secureSalt := wgcrypt.GenRandString()
	securePassword := wgcrypt.HashString((password + secureSalt))

	_, err = prep.Exec(username, securePassword, role, secureSalt)
	if err != nil {
		log.Println("Failed to create account:", err)
	}
}

func DeleteAccount(username string) {
	prep, err := wgdb.Prepare("DELETE FROM account WHERE username = ?;")
	if err != nil {
		log.Println(err)
	}
	defer prep.Close()

	_, err = prep.Exec(username)
	if err != nil {
		log.Println("Failed to delete account data:", err)
	}
}
