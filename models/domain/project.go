package domain

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ProjectCode string `gorm:"primaryKey;type:varchar(15)"`
	ProjectName string `gorm:"type:varchar(100);not null"`
	Deadline    time.Time
	Status      string    `gorm:"type:enum('Complited','On Progress','Delayed');not null"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP();not null"`
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
