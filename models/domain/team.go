package domain

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	TeamCode     string         `gorm:"primaryKey;type:varchar(15);not null"`
	TeamName     string         `gorm:"type:varchar(100);not null"`
	Leader       string         `gorm:"type:varchar(15);not null"`
	ProjectBased bool           `gorm:"default:0;not null"`
	Level        int            `gorm:"type:integer(100);not null"`
	Status       string         `gorm:"type:enum('On Progress','Delayed','Done');not null"`
	CreatedAt    time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP();not null"`
	UpdatedAt    time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP();not null"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
