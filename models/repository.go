package models

import "database/sql"

// Repository is the interface that wraps the basic methods for a repository
type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
	GetBreedByName(string) (*DogBreed, error)
	GetDogOfMonthByID(id int) (*DogOfMonth, error)
}

type mysqlRepository struct {
	DB *sql.DB
}

// NewMysqlRepository creates a new instance of the mysqlRepository type
func NewMysqlRepository(conn *sql.DB) Repository {
	return &mysqlRepository{
		DB: conn,
	}
}

type testRepository struct {
	DB *sql.DB
}

// NewTestRepository creates a new instance of the mysqlRepository type
func NewTestRepository(conn *sql.DB) Repository {
	return &testRepository{
		DB: nil,
	}
}
