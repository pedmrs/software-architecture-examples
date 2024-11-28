package ports

import (
	"hexagonal/internal/domain"
)

type AddressService interface {
	CreateAddress(address domain.Address) (domain.Address, error)
	GetAddress(id int) (domain.Address, error)
	UpdateAddress(id int, updatedAddress domain.Address) (domain.Address, error)
	DeleteAddress(id int) error
}
