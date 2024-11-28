package main

import (
	httpAdapter "hexagonal/internal/adapters/http"
	"hexagonal/internal/adapters/repository"
	"hexagonal/internal/application"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Repositories (Adapters)
	personRepo := repository.NewInMemoryPersonRepository()
	addressRepo := repository.NewInMemoryAddressRepository()

	// Services (Application Logic)
	personService := application.NewPersonService(personRepo)
	addressService := application.NewAddressService(addressRepo)

	// HTTP Handlers (Ports)
	personHandler := httpAdapter.NewPersonHandler(personService)
	addressHandler := httpAdapter.NewAddressHandler(addressService)

	// Set up router
	r := mux.NewRouter()

	// Person routes
	r.HandleFunc("/persons", personHandler.CreatePerson).Methods("POST")
	r.HandleFunc("/persons/{id:[0-9]+}", personHandler.GetPerson).Methods("GET")
	r.HandleFunc("/persons/{id:[0-9]+}", personHandler.UpdatePerson).Methods("PUT")
	r.HandleFunc("/persons/{id:[0-9]+}", personHandler.DeletePerson).Methods("DELETE")

	// Address routes
	r.HandleFunc("/addresses", addressHandler.CreateAddress).Methods("POST")
	r.HandleFunc("/addresses/{id:[0-9]+}", addressHandler.GetAddress).Methods("GET")
	r.HandleFunc("/addresses/{id:[0-9]+}", addressHandler.UpdateAddress).Methods("PUT")
	r.HandleFunc("/addresses/{id:[0-9]+}", addressHandler.DeleteAddress).Methods("DELETE")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
