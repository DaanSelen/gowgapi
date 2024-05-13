package wgrest

import (
	"encoding/json"
	"gowgapi/wgauth"
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

	if firstRun || (wgauth.AuthCredentials(accData.Auth.Username, accData.Auth.Password) && wgauth.AuthAdminRole(accData.Auth.Username)) {
		if accData.Account.Role == "administrator" || accData.Account.Role == "user" {
			ok := wgsqlite.SaveAccount(accData.Account.Username, accData.Account.Password, accData.Account.Role)
			if ok {
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
	firstRun = wgsqlite.CheckEmptyAccountTable()
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

	if wgauth.AuthCredentials(ifaceData.Auth.Username, ifaceData.Auth.Password) {

		//INTERACT WITH SQLITE FOR INTERFACE CONFIG GENERATION

		setCreated(w)
	} else {
		setUnauthorized(w)
	}
}
