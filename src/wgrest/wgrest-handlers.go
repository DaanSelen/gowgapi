package wgrest

import (
	"encoding/json"
	"gowgapi/wgauth"
	"gowgapi/wgparser"
	"gowgapi/wgsqlite"
	"log"
	"net/http"
	"strings"
)

func rootEndpoint(w http.ResponseWriter, r *http.Request) { // BASIC ROOT ENDPOINT RESPONSE
	w.Header().Set("Content-Type", "application/json")
	setOkay(w)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var accData NewAccountBody
	err := json.NewDecoder(r.Body).Decode(&accData)
	if err != nil {
		log.Println(err)
		return
	}

	accData.Account.Role = strings.ToLower(accData.Account.Role) // Convert string entry to lowercase to minimize errors.

	firstRun = wgsqlite.CheckEmptyAccountTable()
	if firstRun || (wgauth.AuthCredentials(accData.Auth.Username, accData.Auth.Password) && wgauth.AuthAdminRole(accData.Auth.Username)) {
		if accData.Account.Role == "administrator" || accData.Account.Role == "user" {
			if ok := wgsqlite.SaveAccount(accData.Account.Username, accData.Account.Password, accData.Account.Role); ok {
				setCreated(w)
			} else {
				setDuplicate(w)
			}
		} else {
			setBad(w)
		}
	} else {
		log.Println("Denied access for:", r.RemoteAddr)
		setUnauthorized(w)
	}
}

func removeAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var accData NewAccountBody
	err := json.NewDecoder(r.Body).Decode(&accData)
	if err != nil {
		log.Println(err)
		return
	}

	if wgauth.AuthCredentials(accData.Auth.Username, accData.Auth.Password) && wgauth.AuthAdminRole(accData.Auth.Username) {
		wgsqlite.DeleteAccount(accData.Account.Username)
		setOkay(w)
	} else {
		log.Println("Denied access for", accData.Account.Username, "on", r.RemoteAddr) // Remove line when being spammed.
		setUnauthorized(w)
	}
}

func createInterface(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ifaceData InterfaceBody
	err := json.NewDecoder(r.Body).Decode(&ifaceData)
	if err != nil {
		return
	}

	if wgauth.AuthCredentials(ifaceData.Auth.Username, ifaceData.Auth.Password) && wgauth.AuthAdminRole(ifaceData.Auth.Username) {
		if ok := wgsqlite.SaveInterface(ifaceData.Interface.Name, ifaceData.Interface.Address, ifaceData.Interface.Port, ifaceData.Interface.Out_Interface, ifaceData.Interface.Description); ok {
			setCreated(w)
		} else {
			setDuplicate(w)
		}
	} else {
		setUnauthorized(w)
	}
}

func parseInterface(w http.ResponseWriter, r *http.Request) {
	wgparser.ParseAll()
}
