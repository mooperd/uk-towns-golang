package entities

import (
	"gorm.io/gorm"
	"time"
)

type Town struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(250);"`
	Latitude    float64 `gorm:"type:decimal(10,2);"`
	Longitude   float64 `gorm:"type:decimal(10,2);"`
	CreatedAt   time.Time
}