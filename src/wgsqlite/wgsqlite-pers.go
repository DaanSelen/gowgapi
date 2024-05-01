package wgsqlite

import (
	"gowgapi/wgcrypt"
	"log"
)

func SaveAccount(username, password, role string) {
	prep, _ := wgdb.Prepare("INSERT INTO user (username, password, role, salt) VALUES ('?', '?', '?', '?');")
	defer prep.Close()

	secureSalt := wgcrypt.GenRandString()
	securePassword := wgcrypt.HashString((password + secureSalt))

	_, err := prep.Exec(prep, username, securePassword, role, secureSalt)
	if err != nil {
		log.Println("Failed to create account:", err)
	}
}
