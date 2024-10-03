package service

import (
	"TDD-GoAPI/model"
	"TDD-GoAPI/repository"
	"fmt"
	"reflect"
	"testing"
)

type MockDB struct {
	books []model.Book
}

func NewBookService(mockDB *MockDB) *bookService {
	return &bookService{
		repository: &repository.Repository{
			Book: mockDB,
		},
	}
}

func newMockDB() *MockDB {
	return &MockDB{
		books: []model.Book{{ID: uint(1), Title: "Test 1", Author: "Test 2"},
			{ID: uint(2), Title: "Test 2", Author: "Test 3"},
		},
	}
}

func TestBookService_GetAllBooks(t *testing.T) {
	mockDB := newMockDB()
	service := NewBookService(mockDB)
	got, err := service.GetAllBooks()
	if err != nil {
		t.Fatal(err)
	}
	want := mockDB.books
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBookService_GetBookByID(t *testing.T) {
	t.Run("Valid ID", func(t *testing.T) {
		mockDB := newMockDB()
		service := NewBookService(mockDB)
		got, err := service.GetBookByID(uint(3))
		want := mockDB.books[1]
		if err != nil {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Invalid ID", func(t *testing.T) {
		mockDB := newMockDB()
		service := NewBookService(mockDB)
		_, err := service.GetBookByID(uint(2))
		wantErr := fmt.Errorf("book with ID %d not found", uint(3))
		if err == nil {
			t.Errorf("expected an error but got nil")
		} else if err.Error() != wantErr.Error() {
			t.Errorf("got %v, want %v", wantErr, err)
		}
	})
}

func TestBookService_CreateBook(t *testing.T) {
	mockDB := &MockDB{
		books: []model.Book{{}},
	}
	service := NewBookService(mockDB)
	newBook := model.Book{
		ID:     5,
		Title:  "Test Title 5",
		Author: "Test Author 5",
	}
	got, err := service.CreateBook(newBook)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if got != newBook {
		t.Errorf("got %v, want %v", got, newBook)
	}
}

func TestBookService_UpdateBook(t *testing.T) {

	mockDB := newMockDB()
	service := NewBookService(mockDB)

	updatedBook := model.Book{
		ID:     2,
		Title:  "Test update Title 2",
		Author: "Test update author 2",
	}
	result, err := service.UpdateBook(updatedBook)
	if err != nil {
		t.Errorf("No book with ID: %v was found", updatedBook.ID)
	}
	if result.Title != updatedBook.Title || result.Author != updatedBook.Author {
		t.Errorf("Expected updated book, but got: %+v", result)
	}
}

func TestBookService_DeleteBook(t *testing.T) {
	mockDB := newMockDB()
	service := NewBookService(mockDB)
	idToDelete := uint(1)
	err := service.DeleteBook(idToDelete)
	if err != nil {
		t.Errorf("No book with ID: %v was found", idToDelete)
	}
}

func (m *MockDB) Update(book model.Book) (model.Book, error) {
	for i, b := range m.books {
		if b.ID == book.ID {
			m.books[i].Title = book.Title
			m.books[i].Author = book.Author
			return m.books[i], nil
		}
	}
	return model.Book{}, fmt.Errorf("book with ID %d not found", book.ID)
}

func (m *MockDB) GetByID(id uint) (model.Book, error) {
	for _, book := range m.books {
		if book.ID == id {
			return book, nil
		}
	}
	return model.Book{}, fmt.Errorf("book with ID %d not found", id)
}

func (m *MockDB) Create(book model.Book) (model.Book, error) {
	m.books = append(m.books, book)
	return book, nil
}

func (m *MockDB) Migrate() error {
	return nil
}

func (m *MockDB) GetAll() ([]model.Book, error) {
	return m.books, nil
}

func (m *MockDB) Delete(id uint) error {
	for _, b := range m.books {
		if b.ID == id {
			return nil
		}
	}
	return fmt.Errorf("book with ID %d not found", id)
}
