package user

import (
	"log"

	"github.com/google/uuid"
)

//Repository is a Repository handler interface
type Repository interface {
	Create(user *User) error
}

type repo struct {
	log *log.Logger
}

//NewRepo is a repositories handler
func NewRepo(l *log.Logger) Repository {
	return &repo{
		log: l,
	}
}

func (repo *repo) Create(user *User) error {

	user.ID = uuid.New().String()

	repo.log.Println("user created with id: ", user.ID)
	return nil
}
