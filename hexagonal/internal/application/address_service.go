package application

import (
	"errors"
	"hexagonal/internal/domain"
)

type AddressRepository interface {
	Save(address domain.Address) (domain.Address, error)
	FindByID(id int) (domain.Address, error)
	Update(id int, updatedAddress domain.Address) (domain.Address, error)
	Delete(id int) error
}

type AddressService struct {
	repo AddressRepository
}

func NewAddressService(repo AddressRepository) *AddressService {
	return &AddressService{repo: repo}
}

func (s *AddressService) CreateAddress(address domain.Address) (domain.Address, error) {
	return s.repo.Save(address)
}

func (s *AddressService) GetAddress(id int) (domain.Address, error) {
	address, err := s.repo.FindByID(id)
	if err != nil {
		return domain.Address{}, errors.New("address not found")
	}
	return address, nil
}

func (s *AddressService) UpdateAddress(id int, updatedAddress domain.Address) (domain.Address, error) {
	return s.repo.Update(id, updatedAddress)
}

func (s *AddressService) DeleteAddress(id int) error {
	return s.repo.Delete(id)
}
