package enrollment

import (
	"log"

	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(enroll *Enrollment) error
	}

	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepo(db *gorm.DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

func (r *repo) Create(enroll *Enrollment) error {

	if err := r.db.Create(enroll).Error; err != nil {
		r.log.Printf("error: %v", err)
		return err
	}

	r.log.Println("enrollment created with id: ", enroll.ID)
	return nil
}
