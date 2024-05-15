package wgsqlite

import (
	"gowgapi/wgcrypt"
	"gowgapi/wgiface"
	"log"
)

func SaveAccount(username, password, role string) bool {
	if checkDuplicateUser(username) {
		return false
	} else {
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
		return true
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

func SaveInterface(ifaceName, ifaceAddr, ifacePort, ifaceOut, ifaceDesc string) bool {
	if checkDuplicateInterface(ifaceName) || checkDuplicateNetwork(ifaceAddr, ifacePort) {
		return false
	} else {
		prep, err := wgdb.Prepare("INSERT INTO iface (name, addr, port, out_interface, privkey, description) VALUES (?, ?, ?, ?, ?, ?);")
		if err != nil {
			log.Println(err)
		}
		defer prep.Close()

		_, err = prep.Exec(ifaceName, ifaceAddr, ifacePort, ifaceOut, wgiface.GenPrivKey(), ifaceDesc)
		if err != nil {
			log.Println("Failed to create interface:", err)
		}
		return true
	}
}
