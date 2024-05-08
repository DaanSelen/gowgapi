package wgrest

import (
	"encoding/json"
	"gowgapi/wgauth"
	"gowgapi/wgsqlite"
	"log"
	"net/http"
	"strings"
)

func createInterface(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var iFaceData InterfaceBody
	err := json.NewDecoder(r.Body).Decode(&iFaceData)
	if err != nil {
		return
	}

	if wgauth.AuthCred(iFaceData.Username, iFaceData.Password) {
		w.WriteHeader(http.StatusCreated)

		//INTERACT WITH SQLITE FOR INTERFACE CONFIG GENERATION

		json.NewEncoder(w).Encode(InfoBody{
			Code:    "CREATED",
			Message: version,
		})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(InfoBody{
			Code:    "UNAUTHORIZED",
			Message: version,
		})
	}
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
	if wgauth.AuthAddr(r.RemoteAddr) {
		if accData.Role == "administrator" || accData.Role == "user" {
			wgsqlite.SaveAccount(accData.Username, accData.Password, accData.Role)
		}
	} else {
		log.Println("Denied access for:", r.RemoteAddr)
	}
	json.NewEncoder(w).Encode(InfoBody{
		Code:    "CREATED",
		Message: version,
	})
}
