package adapters

import (
	"errors"
	"hexagonal/internal/domain"
)

type InMemoryAddressRepository struct {
	data   map[int]domain.Address
	nextID int
}

func NewInMemoryAddressRepository() *InMemoryAddressRepository {
	return &InMemoryAddressRepository{
		data:   make(map[int]domain.Address),
		nextID: 1,
	}
}

func (r *InMemoryAddressRepository) Save(address domain.Address) (domain.Address, error) {
	address.ID = r.nextID
	r.nextID++
	r.data[address.ID] = address
	return address, nil
}

func (r *InMemoryAddressRepository) FindByID(id int) (domain.Address, error) {
	address, exists := r.data[id]
	if !exists {
		return domain.Address{}, errors.New("address not found")
	}
	return address, nil
}

func (r *InMemoryAddressRepository) Update(id int, updatedAddress domain.Address) (domain.Address, error) {
	_, exists := r.data[id]
	if !exists {
		return domain.Address{}, errors.New("address not found")
	}
	updatedAddress.ID = id
	r.data[id] = updatedAddress
	return updatedAddress, nil
}

func (r *InMemoryAddressRepository) Delete(id int) error {
	_, exists := r.data[id]
	if !exists {
		return errors.New("address not found")
	}
	delete(r.data, id)
	return nil
}
