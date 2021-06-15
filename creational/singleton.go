package creational

import "github.com/google/uuid"

type Singleton struct {

	// The "id" is shared and unique between all the instances
	id uuid.UUID

	// The "no" show the number of instantion and increases on every
	// instantiating the Singleton struct
	no int
}

var singleton *Singleton

func NewSingleton() *Singleton {
	if singleton == nil {
		id := uuid.New()
		singleton = &Singleton{id: id}
	}
	singleton.no++
	return singleton
}

func (s *Singleton) ID() uuid.UUID {
	return s.id
}

func (s *Singleton) NO() int {
	return s.no
}
