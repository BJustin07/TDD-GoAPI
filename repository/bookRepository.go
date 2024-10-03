package repository

import (
	"TDD-GoAPI/model"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll() ([]model.Book, error)
	GetByID(id uint) (model.Book, error)
	Create(model.Book) (model.Book, error)
	Update(book model.Book) (model.Book, error)
	Delete(id uint) error
	Migrate() error
}

type bookRepository struct {
	db *gorm.DB
}

func (br *bookRepository) Migrate() error {
	if err := br.db.AutoMigrate(&model.Book{}); err != nil {
		return err
	}
	return nil
}

func (br *bookRepository) GetAll() ([]model.Book, error) {
	var books []model.Book
	result := br.db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
func (br *bookRepository) GetByID(id uint) (model.Book, error) {
	var book model.Book
	result := br.db.First(&book, id)
	if result.Error != nil {
		return book, result.Error
	}
	return book, nil
}

func (br *bookRepository) Create(book model.Book) (model.Book, error) {
	result := br.db.Create(&book)
	if result.Error != nil {
		return book, result.Error
	}
	return book, nil
}

func (br *bookRepository) Update(book model.Book) (model.Book, error) {
	result := br.db.Save(&book)
	if result.Error != nil {
		return book, result.Error
	}
	return book, nil
}

func (br *bookRepository) Delete(id uint) error {
	result := br.db.Delete(&model.Book{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
