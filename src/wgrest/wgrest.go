package wgrest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitFrontend() {
	log.Println("GoWGApi Ready")
	wgapi := mux.NewRouter().StrictSlash(true)

	wgapi.HandleFunc("/", rootEndpoint).Methods("GET")
	wgapi.HandleFunc("/interface/new", createInterface).Methods("POST")

	http.ListenAndServe((":4080"), wgapi)
}

func rootEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(infoBody{
		Code:    "OK",
		Message: "GoWGApi, V0.0.1",
	})
}

func createInterface(w http.ResponseWriter, r *http.Request) {
	var body interfaceBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return
	}

	log.Println(body)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(infoBody{
		Code:    "OK",
		Message: "GoWGApi, V0.0.1",
	})
}
