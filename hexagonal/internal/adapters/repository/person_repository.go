package repository

import (
	"errors"
	"hexagonal/internal/domain"
)

type InMemoryPersonRepository struct {
	data   map[int]domain.Person
	nextID int
}

func NewInMemoryPersonRepository() *InMemoryPersonRepository {
	return &InMemoryPersonRepository{
		data:   make(map[int]domain.Person),
		nextID: 1,
	}
}

func (r *InMemoryPersonRepository) Save(person domain.Person) (domain.Person, error) {
	person.ID = r.nextID
	r.nextID++
	r.data[person.ID] = person
	return person, nil
}

func (r *InMemoryPersonRepository) FindByID(id int) (domain.Person, error) {
	person, exists := r.data[id]
	if !exists {
		return domain.Person{}, errors.New("person not found")
	}
	return person, nil
}

func (r *InMemoryPersonRepository) Update(id int, updatedPerson domain.Person) (domain.Person, error) {
	_, exists := r.data[id]
	if !exists {
		return domain.Person{}, errors.New("person not found")
	}
	updatedPerson.ID = id
	r.data[id] = updatedPerson
	return updatedPerson, nil
}

func (r *InMemoryPersonRepository) Delete(id int) error {
	_, exists := r.data[id]
	if !exists {
		return errors.New("person not found")
	}
	delete(r.data, id)
	return nil
}
