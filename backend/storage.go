package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreatePetrolModel(*PetrolModel) error
}

type PostgresStorage struct {
	db *sql.DB
}

func (s *PostgresStorage) CreatePetrolModel(p *PetrolModel) error {
	return nil
}
