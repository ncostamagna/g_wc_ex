package enrollment

import (
	"errors"
	"log"

	"github.com/ncostamagna/g_wc_ex/internal/course"
	"github.com/ncostamagna/g_wc_ex/internal/domain"
	"github.com/ncostamagna/g_wc_ex/internal/user"
)

type (
	Filters struct {
		UserID   string
		CourseID string
	}

	Service interface {
		Create(userID, courseID string) (*domain.Enrollment, error)
	}

	service struct {
		log        *log.Logger
		userRepo   user.Service
		courseRepo course.Service
		repo       Repository
	}
)

func NewService(l *log.Logger, userRepo user.Service, courseRepo course.Service, repo Repository) Service {
	return &service{
		log:        l,
		userRepo:   userRepo,
		courseRepo: courseRepo,
		repo:       repo,
	}
}

func (s service) Create(userID, courseID string) (*domain.Enrollment, error) {

	enroll := &domain.Enrollment{
		UserID:   userID,
		CourseID: courseID,
		Status:   "P",
	}

	if _, err := s.userRepo.Get(enroll.UserID); err != nil {
		return nil, errors.New("user id doesn't exists")
	}

	if _, err := s.courseRepo.Get(enroll.CourseID); err != nil {
		return nil, errors.New("course id doesn't exists")
	}

	if err := s.repo.Create(enroll); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return enroll, nil
}
