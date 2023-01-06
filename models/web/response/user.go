package response

import (
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	Id        int            `json:"user_id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
