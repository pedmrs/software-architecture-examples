package repository

import (
	"layered/internal/service"
	"sync"
)

type inMemoryPersonRepository struct {
	persons map[int]service.Person
	mu      sync.Mutex
	counter int
}

func NewInMemoryPersonRepository() service.PersonRepository {
	return &inMemoryPersonRepository{
		persons: make(map[int]service.Person),
		counter: 0,
	}
}

func (r *inMemoryPersonRepository) Create(person service.Person) (service.Person, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.counter++
	person.ID = r.counter
	r.persons[person.ID] = person
	return person, nil
}

func (r *inMemoryPersonRepository) FindByID(id int) (service.Person, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	person, found := r.persons[id]
	return person, found
}

func (r *inMemoryPersonRepository) Update(id int, person service.Person) (service.Person, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	person.ID = id
	r.persons[id] = person
	return person, nil
}

func (r *inMemoryPersonRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.persons, id)
	return nil
}
