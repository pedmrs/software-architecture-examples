package ports

import (
	"encoding/json"
	"hexagonal/internal/domain"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PersonService interface {
	CreatePerson(person domain.Person) (domain.Person, error)
	GetPerson(id int) (domain.Person, error)
	UpdatePerson(id int, updatedPerson domain.Person) (domain.Person, error)
	DeletePerson(id int) error
}

type PersonHandler struct {
	service PersonService
}

func NewPersonHandler(service PersonService) *PersonHandler {
	return &PersonHandler{service: service}
}

func (h *PersonHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person domain.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdPerson, err := h.service.CreatePerson(person)
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
	person, err := h.service.GetPerson(id)
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
	var person domain.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedPerson, err := h.service.UpdatePerson(id, person)
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
	err = h.service.DeletePerson(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}