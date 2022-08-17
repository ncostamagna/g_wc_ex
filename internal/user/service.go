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
	log *log.Logger
}

//NewService is a service handler
func NewService(l *log.Logger) Service {
	return &service{
		log: l,
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
	log.Println("Create user service")
	return user, nil
}
