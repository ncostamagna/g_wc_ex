package user

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Repository is a Repository handler interface
type Repository interface {
	Create(user *User) error
}

type repo struct {
	db  *gorm.DB
	log *log.Logger
}

//NewRepo is a repositories handler
func NewRepo(db *gorm.DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

func (repo *repo) Create(user *User) error {

	user.ID = uuid.New().String()

	repo.log.Println("user created with id: ", user.ID)
	return nil
}
