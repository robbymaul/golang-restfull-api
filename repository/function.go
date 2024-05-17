package repository

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func RepositoryDB(db *gorm.DB) *Repository {
	return &Repository{db}
}
