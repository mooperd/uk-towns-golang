package entities

import (
	"time"
)

type Journey struct {
	Name        string `gorm:"type:varchar(250);"`
	FromTown    Town   `gorm:"ForeignKey:FromTownID;AssociationForeignKey:Id;not null"`
	FromTownID  uint
	ToTown      Town `gorm:"ForeignKey:ToTownID;AssociationForeignKey:Id;not null"`
	ToTownID    uint
	Distance	float64
	CreatedAt   time.Time
}