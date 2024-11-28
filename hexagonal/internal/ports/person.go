package ports

import (
	"hexagonal/internal/domain"
)

type PersonService interface {
	CreatePerson(person domain.Person) (domain.Person, error)
	GetPerson(id int) (domain.Person, error)
	UpdatePerson(id int, updatedPerson domain.Person) (domain.Person, error)
	DeletePerson(id int) error
}
