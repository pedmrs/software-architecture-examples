package handler

import (
	"encoding/json"
	"layered/internal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AddressHandler struct {
	addressService service.AddressService
}

func NewAddressHandler(service service.AddressService) *AddressHandler {
	return &AddressHandler{addressService: service}
}

func (h *AddressHandler) CreateAddress(w http.ResponseWriter, r *http.Request) {
	var address service.Address
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdAddress, err := h.addressService.CreateAddress(address)
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
	address, err := h.addressService.GetAddress(id)
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
	var address service.Address
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedAddress, err := h.addressService.UpdateAddress(id, address)
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
	err = h.addressService.DeleteAddress(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
