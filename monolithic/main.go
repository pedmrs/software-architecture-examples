package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

// Entities
type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Address struct {
	ID      int    `json:"id"`
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

// Repositories
var (
	persons          = make(map[int]Person)
	addresses        = make(map[int]Address)
	personMu         sync.Mutex
	addressMu        sync.Mutex
	personIDCounter  = 0
	addressIDCounter = 0
)

// Handlers for Person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	personMu.Lock()
	defer personMu.Unlock()
	personIDCounter++
	person.ID = personIDCounter
	persons[person.ID] = person
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	personMu.Lock()
	person, exists := persons[id]
	personMu.Unlock()
	if !exists {
		http.Error(w, "person not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(person)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	var updatedPerson Person
	if err := json.NewDecoder(r.Body).Decode(&updatedPerson); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	personMu.Lock()
	defer personMu.Unlock()
	_, exists := persons[id]
	if !exists {
		http.Error(w, "person not found", http.StatusNotFound)
		return
	}
	updatedPerson.ID = id
	persons[id] = updatedPerson
	json.NewEncoder(w).Encode(updatedPerson)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	personMu.Lock()
	defer personMu.Unlock()
	if _, exists := persons[id]; !exists {
		http.Error(w, "person not found", http.StatusNotFound)
		return
	}
	delete(persons, id)
	w.WriteHeader(http.StatusNoContent)
}

// Handlers for Address
func CreateAddress(w http.ResponseWriter, r *http.Request) {
	var address Address
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	addressMu.Lock()
	defer addressMu.Unlock()
	addressIDCounter++
	address.ID = addressIDCounter
	addresses[address.ID] = address
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(address)
}

func GetAddress(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	addressMu.Lock()
	address, exists := addresses[id]
	addressMu.Unlock()
	if !exists {
		http.Error(w, "address not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(address)
}

func UpdateAddress(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	var updatedAddress Address
	if err := json.NewDecoder(r.Body).Decode(&updatedAddress); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	addressMu.Lock()
	defer addressMu.Unlock()
	_, exists := addresses[id]
	if !exists {
		http.Error(w, "address not found", http.StatusNotFound)
		return
	}
	updatedAddress.ID = id
	addresses[id] = updatedAddress
	json.NewEncoder(w).Encode(updatedAddress)
}

func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	addressMu.Lock()
	defer addressMu.Unlock()
	if _, exists := addresses[id]; !exists {
		http.Error(w, "address not found", http.StatusNotFound)
		return
	}
	delete(addresses, id)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/persons", CreatePerson).Methods("POST")
	r.HandleFunc("/persons/{id:[0-9]+}", GetPerson).Methods("GET")
	r.HandleFunc("/persons/{id:[0-9]+}", UpdatePerson).Methods("PUT")
	r.HandleFunc("/persons/{id:[0-9]+}", DeletePerson).Methods("DELETE")

	r.HandleFunc("/addresses", CreateAddress).Methods("POST")
	r.HandleFunc("/addresses/{id:[0-9]+}", GetAddress).Methods("GET")
	r.HandleFunc("/addresses/{id:[0-9]+}", UpdateAddress).Methods("PUT")
	r.HandleFunc("/addresses/{id:[0-9]+}", DeleteAddress).Methods("DELETE")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
