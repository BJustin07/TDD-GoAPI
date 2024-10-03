package service

import (
	"TDD-GoAPI/model"
	"TDD-GoAPI/repository"
)

type (
	BookService interface {
		GetAllBooks() ([]model.Book, error)
		GetBookByID(id uint) (model.Book, error)
		CreateBook(book model.Book) (model.Book, error)
		UpdateBook(book model.Book) (model.Book, error)
		DeleteBook(id uint) error
	}
	bookService struct {
		repository *repository.Repository
	}
)

func (s *bookService) GetAllBooks() ([]model.Book, error) {
	result, err := s.repository.Book.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *bookService) GetBookByID(id uint) (model.Book, error) {
	result, err := s.repository.Book.GetByID(id)
	if err != nil {
		return model.Book{}, err
	}
	return result, nil
}

func (s *bookService) CreateBook(book model.Book) (model.Book, error) {
	result, err := s.repository.Book.Create(book)
	if err != nil {
		return model.Book{}, err
	}
	return result, nil
}

func (s *bookService) UpdateBook(book model.Book) (model.Book, error) {
	var originalBook model.Book
	originalBook, err := s.repository.Book.GetByID(book.ID)
	if err != nil {
		return model.Book{}, err
	}
	originalBook.Title = book.Title
	originalBook.Author = book.Author
	_, err = s.repository.Book.Update(originalBook)
	if err != nil {
		return model.Book{}, err
	}
	return originalBook, nil
}

func (s *bookService) DeleteBook(id uint) error {
	err := s.repository.Book.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
