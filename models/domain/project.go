package domain

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ProjectCode string         
	ProjectName string         
	Deadline    string     
	Status      string         
	CreatedAt   time.Time      
	UpdatedAt   time.Time      
	DeletedAt   gorm.DeletedAt 
}

func (e *Project) TabelName() string {
	return "project"
}
