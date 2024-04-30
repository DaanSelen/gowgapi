package wgrest

// Package for the REST API section of the application.

import (
	"encoding/json"
	"gowgapi/wgauth"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	version string = "GoWGAPI: 0.0.2"
)

func InitFrontend() {
	log.Println("GoWGAPI Ready")
	wgapi := mux.NewRouter().StrictSlash(true)

	wgapi.HandleFunc("/", rootEndpoint).Methods("GET")
	wgapi.HandleFunc("/iface/new", createInterface).Methods("POST")

	http.ListenAndServe((":4080"), wgapi)
}

func rootEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(infoBody{
		Code:    "OK",
		Message: "GoWGAPI, V0.0.1",
	})
}

func createInterface(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var iFaceBody interfaceBody
	err := json.NewDecoder(r.Body).Decode(&iFaceBody)
	if err != nil {
		return
	}

	if wgauth.Authenticate(iFaceBody.Username, iFaceBody.Password) {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(infoBody{
			Code:    "CREATED",
			Message: version,
		})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(infoBody{
			Code:    "UNAUTHORIZED",
			Message: version,
		})
	}
}
