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
	var accData AccountBody
	err := json.NewDecoder(r.Body).Decode(&accData)
	if err != nil {
		log.Println(err)
		return
	}

	accData.Role = strings.ToLower(accData.Role) // Convert string entry to lowercase to minimize errors.
	if wgauth.AuthLocalAddr(r.RemoteAddr) {
		if accData.Role == "administrator" || accData.Role == "user" {
			wgsqlite.SaveAccount(accData.Username, accData.Password, accData.Role)
			setCreated(w)
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
	var accData AccountBody
	err := json.NewDecoder(r.Body).Decode(&accData)
	if err != nil {
		log.Println(err)
		return
	}

	if wgauth.AuthCredentials(accData.Username, accData.Password) && wgauth.AuthAdminRole(accData.Username) {
		wgsqlite.DeleteAccount(accData.Username)
		setOkay(w)
	} else {
		log.Println("Denied access for", accData.Username, "on", r.RemoteAddr) // Remove line when being spammed.
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

	if wgauth.AuthCredentials(ifaceData.Username, ifaceData.Password) {

		//INTERACT WITH SQLITE FOR INTERFACE CONFIG GENERATION

		setCreated(w)
	} else {
		setUnauthorized(w)
	}
}
