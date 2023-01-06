package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (e *User) TableName() string {
	return "user"
}
