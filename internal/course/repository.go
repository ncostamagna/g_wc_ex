package course

import (
	"log"

	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(course *Course) error
	}

	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

//NewRepo is a repositories handler
func NewRepo(db *gorm.DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

func (r *repo) Create(course *Course) error {

	if err := r.db.Create(course).Error; err != nil {
		r.log.Printf("error: %v", err)
		return err
	}

	r.log.Println("course created with id: ", course.ID)
	return nil
}
