package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ionutdejeu/restapi/cmd/locahost"
)

 
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", CreateEvent_Handler).Methods("POST")
	//router.HandleFunc("/events", getAllEvents).Methods("GET")
	//router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	//router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	//router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}