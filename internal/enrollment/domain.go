package enrollment

import (
	"time"

	"github.com/google/uuid"
	"github.com/ncostamagna/g_wc_ex/internal/course"
	"github.com/ncostamagna/g_wc_ex/internal/user"
	"gorm.io/gorm"
)

type Enrollment struct {
	ID        string         `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	UserID    string         `json:"user_id,omitempty" gorm:"type:char(36)"`
	User      *user.User     `json:"user,omitempty"`
	CourseID  string         `json:"course_id" gorm:"type:char(36);not null"`
	Course    *course.Course `json:"course,omitempty"`
	Status    string         `json:"status" gorm:"type:char(2)"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
}

func (c *Enrollment) BeforeCreate(tx *gorm.DB) (err error) {

	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return
}
