package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ionutdejeu/golang-hex-architecture-mongo-db-graph-storage/cmd/localhost"
)

 
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/event",localhost.CreateEvent_Handler).Methods("POST")
	router.HandleFunc("/events",localhost.GetAllEvents_Handler).Methods("GET")
	router.HandleFunc("/events/{id}", localhost.GetOneEvent_Handler).Methods("GET")
	router.HandleFunc("/events/{id}", localhost.UpdateEvent_Handler).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8081", router))
}