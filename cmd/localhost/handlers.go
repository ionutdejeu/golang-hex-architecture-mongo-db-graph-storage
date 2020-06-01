package localhost

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

func CreateEvent_Handler(w http.ResponseWriter, r *http.Request){
	var newEvent Event 
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{ 
		fmt.Println("Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody,&newEvent)
	Events = append(Events,newEvent)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEvent)
}

func GetOneEvent_Handler(w http.ResponseWriter, r *http.Request){
	 
	eventID := mux.Vars(r)["id"]
	// Get the details from an existing event
	// Use the blank identifier to avoid creating a value that will not be used
	for _, singleEvent := range Events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func GetAllEvents_Handler(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(Events)
}

func UpdateEvent_Handler(w http.ResponseWriter, r *http.Request){
	eventID := mux.Vars(r)["id"]
	var updatedEvent Event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error:",err.Error())
	}
	json.Unmarshal(reqBody,&updatedEvent)

	for i,singleEvent := range Events{ 
		if singleEvent.ID == eventID{
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			Events[i] = singleEvent
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}


