package application

import (
	"errors"
	"hexagonal/internal/domain"
)

type PersonRepository interface {
	Save(person domain.Person) (domain.Person, error)
	FindByID(id int) (domain.Person, error)
	Update(id int, updatedPerson domain.Person) (domain.Person, error)
	Delete(id int) error
}

type PersonService struct {
	repo PersonRepository
}

func NewPersonService(repo PersonRepository) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) CreatePerson(person domain.Person) (domain.Person, error) {
	return s.repo.Save(person)
}

func (s *PersonService) GetPerson(id int) (domain.Person, error) {
	person, err := s.repo.FindByID(id)
	if err != nil {
		return domain.Person{}, errors.New("person not found")
	}
	return person, nil
}

func (s *PersonService) UpdatePerson(id int, updatedPerson domain.Person) (domain.Person, error) {
	return s.repo.Update(id, updatedPerson)
}

func (s *PersonService) DeletePerson(id int) error {
	return s.repo.Delete(id)
}
