package wgauth

import (
	"gowgapi/wgcrypt"
	"gowgapi/wgsqlite"
)

func Authenticate(username, password string) bool {
	result := wgsqlite.QueryUser(username)

	wgcrypt.HashString(result.Password)
	return true
}
