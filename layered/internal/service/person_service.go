package service

import "errors"

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PersonService interface {
	CreatePerson(person Person) (Person, error)
	GetPerson(id int) (Person, error)
	UpdatePerson(id int, updatedPerson Person) (Person, error)
	DeletePerson(id int) error
}

type personService struct {
	repo PersonRepository
}

func NewPersonService(repo PersonRepository) PersonService {
	return &personService{repo: repo}
}

func (s *personService) CreatePerson(person Person) (Person, error) {
	return s.repo.Create(person)
}

func (s *personService) GetPerson(id int) (Person, error) {
	person, found := s.repo.FindByID(id)
	if !found {
		return Person{}, errors.New("person not found")
	}
	return person, nil
}

func (s *personService) UpdatePerson(id int, updatedPerson Person) (Person, error) {
	_, found := s.repo.FindByID(id)
	if !found {
		return Person{}, errors.New("person not found")
	}
	return s.repo.Update(id, updatedPerson)
}

func (s *personService) DeletePerson(id int) error {
	_, found := s.repo.FindByID(id)
	if !found {
		return errors.New("person not found")
	}
	return s.repo.Delete(id)
}

type PersonRepository interface {
	Create(person Person) (Person, error)
	FindByID(id int) (Person, bool)
	Update(id int, person Person) (Person, error)
	Delete(id int) error
}
