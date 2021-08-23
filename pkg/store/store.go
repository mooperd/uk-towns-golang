package store

import (
	"gorm.io/gorm"
	"uk-towns/pkg/entities"
)

// keep it private by using lower case name.
type store struct {
	db *gorm.DB
}

type Storage interface {
	CreateTown(town *entities.Town) error
	CreateJourney(journey *entities.Journey) error
	GetJourney() (*entities.Journey, error )
}

func NewStorage(db *gorm.DB) Storage {
	return &store{ db:db }
}