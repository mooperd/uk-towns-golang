package store

import "uk-towns/pkg/entities"

func(s *store) CreateTown(town *entities.Town) error {
	return s.db.Create(&town).Error
}