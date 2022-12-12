package response

import (
	"time"

	"gorm.io/gorm"
)

type ProjectResponse struct {
	ProjectCode string         `json:"project_code"`
	ProjectName string         `json:"project_name"`
	Deadline    string     `json:"deadline"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
