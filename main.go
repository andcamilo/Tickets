package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)

// Types
type ticket struct {
	ID      int    `json:"ID"`
	User    string `json:"User"`
	Status  string `json:"Status"`
	Create  time.Time
	Update  time.Time
}

type alltickets []ticket

// Persistence
var tickets = alltickets{
	{
		ID:      1,
		User:    "Anderson",
		Status: "Abierto",
		Create: time.Now(),
		Update: time.Now(),
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API de Tickets!")
}

func createticket(w http.ResponseWriter, r *http.Request) {
	var newticket ticket
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid ticket Data")
	}

	json.Unmarshal(reqBody, &newticket)
	newticket.ID = len(tickets) + 1
	newticket.Create = time.Now()
	newticket.Update = time.Now()
	tickets = append(tickets, newticket)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newticket)

}

func gettickets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

func getOneticket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}

	for _, ticket := range tickets {
		if ticket.ID == ticketID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ticket)
		}
	}
}

func updateticket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketID, err := strconv.Atoi(vars["id"])
	var updatedticket ticket

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data")
	}
	json.Unmarshal(reqBody, &updatedticket)

	for i, t := range tickets {
		if t.ID == ticketID {
		

			updatedticket.ID = t.ID
			updatedticket.Create = t.Create
			updatedticket.Update = time.Now()
			tickets = append(tickets, updatedticket)

			tickets = append(tickets[:i], tickets[i+1:]...)

			// w.Header().Set("Content-Type", "application/json")
			// json.NewEncoder(w).Encode(updatedticket)
			fmt.Fprintf(w, "The ticket with ID %v has been updated successfully", ticketID)
		}
	}

}

func deleteticket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid User ID")
		return
	}

	for i, t := range tickets {
		if t.ID == ticketID {
			tickets = append(tickets[:i], tickets[i+1:]...)
			fmt.Fprintf(w, "The ticket with ID %v has been remove successfully", ticketID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tickets/create", createticket).Methods("POST")
	router.HandleFunc("/tickets", gettickets).Methods("GET")
	router.HandleFunc("/tickets/{id}", getOneticket).Methods("GET")
	router.HandleFunc("/tickets/{id}", deleteticket).Methods("DELETE")
	router.HandleFunc("/tickets/{id}", updateticket).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}
