package wgrest

// Package for the REST API section of the application.

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

const (
	version  string = "GoWGAPI: 0.0.2"
	certFile string = "./certificate/gowgapi.crt"
	keyFile  string = "./certificate/gowgapi.key"
)

func InitFrontend(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	wgapi := mux.NewRouter().StrictSlash(true)

	secureWeb := &http.Server{
		Addr:     "0.0.0.0:4080", // Specify the desired HTTPS port
		Handler:  wgapi,
		ErrorLog: log.New(io.Discard, "", 0), // THIS IS DONE TO NOT RECEIVE CLIENT HTTP ERRORS. TO DEBUG, REMOVE THIS LINE OR CREATE A VALID LOGGER
	}

	wgapi.HandleFunc("/", rootEndpoint).Methods("GET")

	wgapi.HandleFunc("/account/create", createAccount).Methods("POST")
	wgapi.HandleFunc("/account/delete", removeAccount).Methods("DELETE")

	wgapi.HandleFunc("/iface/create", createInterface).Methods("POST")
	wgapi.HandleFunc("/iface/delete", createInterface).Methods("DELETE")

	err := secureWeb.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatal("Failed to launch REST HTTP API:", err)
	}
}

func setOkay(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(InfoBody{
		Code:    "OK REQUEST",
		Message: version,
	})
}

func setCreated(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(InfoBody{
		Code:    "CREATED",
		Message: version,
	})
}

func setUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(InfoBody{
		Code:    "UNAUTHORIZED REQUEST",
		Message: version,
	})
}

func setBad(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(InfoBody{
		Code:    "BAD REQUEST",
		Message: version,
	})
}
