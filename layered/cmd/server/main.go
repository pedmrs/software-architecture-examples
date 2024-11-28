package main

import (
	"log"
	"net/http"

	"layered/internal/handler"
	"layered/internal/repository"
	"layered/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	// Repositories
	personRepo := repository.NewInMemoryPersonRepository()
	addressRepo := repository.NewInMemoryAddressRepository()

	// Services
	personService := service.NewPersonService(personRepo)
	addressService := service.NewAddressService(addressRepo)

	// Handlers
	personHandler := handler.NewPersonHandler(personService)
	addressHandler := handler.NewAddressHandler(addressService)

	// Setup router
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

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
