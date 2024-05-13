package wgauth

// Package for authenticating user from the application.

import (
	"gowgapi/wgcrypt"
	"gowgapi/wgsqlite"
)

func AuthCredentials(username, password string) bool {
	result := wgsqlite.QueryUser(username)
	if len(result.Username) == 0 || len(result.Password) == 0 {
		return false
	}

	candidatePass := wgcrypt.HashString(password + result.Salt)
	return candidatePass == result.Password
}

func AuthAdminRole(username string) bool {
	result := wgsqlite.QueryUser(username)
	if len(result.Username) == 0 {
		return false
	}

	return result.Role == "administrator"
}
