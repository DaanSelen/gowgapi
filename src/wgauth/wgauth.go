package wgauth

import (
	"gowgapi/wgcrypt"
	"gowgapi/wgsqlite"
	"log"
)

func Authenticate(username, password string) bool {
	result := wgsqlite.QueryUser(username)

	candidatePass := wgcrypt.HashString(password)
	returnedPass := wgcrypt.HashString(result.Password)
	log.Println(candidatePass)
	log.Println(returnedPass)
	return true
}
