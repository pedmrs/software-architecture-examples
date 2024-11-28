package service

import "errors"

type Address struct {
	ID      int    `json:"id"`
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

type AddressService interface {
	CreateAddress(address Address) (Address, error)
	GetAddress(id int) (Address, error)
	UpdateAddress(id int, updatedAddress Address) (Address, error)
	DeleteAddress(id int) error
}

type addressService struct {
	repo AddressRepository
}

func NewAddressService(repo AddressRepository) AddressService {
	return &addressService{repo: repo}
}

func (s *addressService) CreateAddress(address Address) (Address, error) {
	return s.repo.Create(address)
}

func (s *addressService) GetAddress(id int) (Address, error) {
	address, found := s.repo.FindByID(id)
	if !found {
		return Address{}, errors.New("address not found")
	}
	return address, nil
}

func (s *addressService) UpdateAddress(id int, updatedAddress Address) (Address, error) {
	_, found := s.repo.FindByID(id)
	if !found {
		return Address{}, errors.New("address not found")
	}
	return s.repo.Update(id, updatedAddress)
}

func (s *addressService) DeleteAddress(id int) error {
	_, found := s.repo.FindByID(id)
	if !found {
		return errors.New("address not found")
	}
	return s.repo.Delete(id)
}

type AddressRepository interface {
	Create(address Address) (Address, error)
	FindByID(id int) (Address, bool)
	Update(id int, address Address) (Address, error)
	Delete(id int) error
}
