package user

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Repository is a Repository handler interface
type Repository interface {
	Create(user *User) error
	GetAll() ([]User, error)
	Get(id string) (*User, error)
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

func (r *repo) Create(user *User) error {

	user.ID = uuid.New().String()

	if err := r.db.Create(user).Error; err != nil {
		r.log.Printf("error: %v", err)
		return err
	}

	r.log.Println("user created with id: ", user.ID)
	return nil
}

func (r *repo) GetAll() ([]User, error) {
	var u []User
	result := r.db.Model(&u).Order("created_at desc").Find(&u)

	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func (r *repo) Get(id string) (*User, error) {
	user := User{ID: id}
	result := r.db.First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
