package store

import (
	"uk-towns/pkg/entities"
)


func(s *store) CreateJourney(journey *entities.Journey) error {
	return s.db.Create(&journey).Error
}

func (s *store) GetJourney() (*entities.Journey, error) {
	journey := new(entities.Journey)
	err := s.db.Preload("FromTown").Preload("ToTown").First(journey).Error
	if err != nil {
		return nil, err
	}
	return journey, nil
}