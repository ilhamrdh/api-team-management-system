package web

import (
	"time"

	"gorm.io/gorm"
)

type TeamResponse struct {
	TeamCode     string         `json:"team_code"`
	TeamName     string         `json:"team_name"`
	Leader       string         `json:"leader"`
	ProjectBased bool           `json:"project_based"`
	Level        int            `json:"level"`
	Status       string         `json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
