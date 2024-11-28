package ports

import (
	"encoding/json"
	"hexagonal/internal/domain"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AddressService interface {
	CreateAddress(address domain.Address) (domain.Address, error)
	GetAddress(id int) (domain.Address, error)
	UpdateAddress(id int, updatedAddress domain.Address) (domain.Address, error)
	DeleteAddress(id int) error
}

type AddressHandler struct {
	service AddressService
}

func NewAddressHandler(service AddressService) *AddressHandler {
	return &AddressHandler{service: service}
}

func (h *AddressHandler) CreateAddress(w http.ResponseWriter, r *http.Request) {
	var address domain.Address
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdAddress, err := h.service.CreateAddress(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdAddress)
}

func (h *AddressHandler) GetAddress(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	address, err := h.service.GetAddress(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(address)
}

func (h *AddressHandler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	var address domain.Address
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedAddress, err := h.service.UpdateAddress(id, address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedAddress)
}

func (h *AddressHandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	err = h.service.DeleteAddress(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
