package wgauth

// Package for authenticating user from the application.

import (
	"gowgapi/wgcrypt"
	"gowgapi/wgsqlite"
	"log"
	"net"
)

func AuthCred(username, password string) bool {
	result := wgsqlite.QueryUser(username)
	if len(result.Username) == 0 || len(result.Password) == 0 {
		return false
	}

	candidatePass := wgcrypt.HashString(password)
	returnedPass := wgcrypt.HashString(result.Password)
	return candidatePass == returnedPass
}

func AuthAddr(remoteAddr string) bool {
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		log.Println("Error reading remote address.")
	}

	if host == "127.0.0.1" || host == "::1" {
		return true
	} else {
		return false
	}
}
