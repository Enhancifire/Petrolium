package main

import (
	"database/sql"
	"github.com/Enhancifire/Petrolium/backend/models"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreatePetrolModel(*models.PetrolModel) error
}

type PostgresStorage struct {
	db *sql.DB
}

func (s *PostgresStorage) CreatePetrolModel(p *models.PetrolModel) error {
	return nil
}
