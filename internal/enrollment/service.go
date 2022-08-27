package enrollment

import (
	"log"
)

type (
	Filters struct {
		UserID   string
		CourseID string
	}

	Service interface {
		Create(userID, courseID string) (*Enrollment, error)
	}

	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(l *log.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}

func (s service) Create(userID, courseID string) (*Enrollment, error) {

	enroll := &Enrollment{
		UserID:   userID,
		CourseID: courseID,
		Status:   "P",
	}

	if err := s.repo.Create(enroll); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return enroll, nil
}
