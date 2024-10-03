package service

import "TDD-GoAPI/repository"

type Services struct {
	Book BookService
}

func New(r *repository.Repository) *Services {
	return &Services{
		Book: &bookService{repository: r},
	}
}
