package repository

import (
	"layered/internal/service"
	"sync"
)

type inMemoryAddressRepository struct {
	addresses map[int]service.Address
	mu        sync.Mutex
	counter   int
}

func NewInMemoryAddressRepository() service.AddressRepository {
	return &inMemoryAddressRepository{
		addresses: make(map[int]service.Address),
		counter:   0,
	}
}

func (r *inMemoryAddressRepository) Create(address service.Address) (service.Address, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.counter++
	address.ID = r.counter
	r.addresses[address.ID] = address
	return address, nil
}

func (r *inMemoryAddressRepository) FindByID(id int) (service.Address, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	address, found := r.addresses[id]
	return address, found
}

func (r *inMemoryAddressRepository) Update(id int, address service.Address) (service.Address, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	address.ID = id
	r.addresses[id] = address
	return address, nil
}

func (r *inMemoryAddressRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.addresses, id)
	return nil
}
