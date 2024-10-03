package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB   *gorm.DB
	Book BookRepository
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		DB:   db,
		Book: &bookRepository{db},
	}
}
