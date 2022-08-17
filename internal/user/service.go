package user

import (
	"log"
)

// primero servicio
// despues mostramos como setar el logger
// despues pasar el objeto User
type Service interface {
	Create(firstName, lastName, email, phone string) (*User, error)
}

type service struct {
	log  *log.Logger
	repo Repository
}

//NewService is a service handler
func NewService(l *log.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}

//Create service
func (s service) Create(firstName, lastName, email, phone string) (*User, error) {
	user := &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}
	s.log.Println("Create user service")

	if err := s.repo.Create(user); err != nil {
		s.log.Println(err)
	}

	return user, nil
}
