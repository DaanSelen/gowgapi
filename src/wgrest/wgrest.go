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

var ()

func InitFrontend(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	wgapi := mux.NewRouter().StrictSlash(true)

	secureWeb := &http.Server{
		Addr:     "0.0.0.0:4080", // Specify the desired HTTPS port
		Handler:  wgapi,
		ErrorLog: log.New(io.Discard, "", 0), // THIS IS DONE TO NOT RECEIVE CLIENT HTTP ERRORS. TO DEBUG, REMOVE THIS LINE OR CREATE A VALID LOGGER
	}

	wgapi.HandleFunc("/", rootEndpoint).Methods("GET")

	wgapi.HandleFunc("/account/new", createAccount).Methods("POST")
	wgapi.HandleFunc("/iface/new", createInterface).Methods("POST")

	err := secureWeb.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatal("Failed to launch REST HTTP API:", err)
	}
}

func rootEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(InfoBody{
		Code:    "OK",
		Message: "GoWGAPI, V0.0.1",
	})
}
