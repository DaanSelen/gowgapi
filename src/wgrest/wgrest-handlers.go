package wgrest

import (
	"encoding/json"
	"gowgapi/wgauth"
	"gowgapi/wgsqlite"
	"net/http"
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
		return
	}

	if wgauth.AuthAddr(r.RemoteAddr) {
		wgsqlite.SaveAccount(accData.Username, accData.Password, accData.Role)
	}
}
