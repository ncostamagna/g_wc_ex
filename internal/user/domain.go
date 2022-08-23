package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	CreatedAt *time.Time     `json:"-"`
	UpdateAt  *time.Time     `json:"-"`
	Deleted   gorm.DeletedAt `json:"-"`
}
