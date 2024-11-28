package handler

import (
	"encoding/json"
	"layered/internal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PersonHandler struct {
	personService service.PersonService
}

func NewPersonHandler(service service.PersonService) *PersonHandler {
	return &PersonHandler{personService: service}
}

func (h *PersonHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person service.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdPerson, err := h.personService.CreatePerson(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPerson)
}

func (h *PersonHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	person, err := h.personService.GetPerson(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(person)
}

func (h *PersonHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	var person service.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedPerson, err := h.personService.UpdatePerson(id, person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedPerson)
}

func (h *PersonHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	err = h.personService.DeletePerson(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
