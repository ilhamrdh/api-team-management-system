package domain

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	TeamCode     string         
	TeamName     string         
	Leader       string         
	ProjectBased bool           
	Level        int            
	Status       string         
	CreatedAt    time.Time      
	UpdatedAt    time.Time      
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (e *Team) TabelName() string {
	return "team"
}
