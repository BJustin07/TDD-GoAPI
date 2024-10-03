package controller

import (
	"TDD-GoAPI/model"
	"TDD-GoAPI/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type (
	BookController interface {
		GetBook(c echo.Context) error
		GetAllBooks(c echo.Context) error
		//TestMessage(c echo.Context) error
		CreateBook(c echo.Context) error
		UpdateBook(c echo.Context) error
		DeleteBook(c echo.Context) error
	}
	bookController struct {
		service.BookService
	}
)

func (b *bookController) GetAllBooks(c echo.Context) error {
	var books []model.Book
	books, err := b.BookService.GetAllBooks()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, books)
}

func (b *bookController) GetBook(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	result, err := b.BookService.GetBookByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (b *bookController) CreateBook(c echo.Context) error {
	var book model.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	book, err := b.BookService.CreateBook(book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, book)
}

func (b *bookController) UpdateBook(c echo.Context) error {
	idParam := c.Param("id")
	id, errID := strconv.ParseUint(idParam, 10, 64)
	if errID != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}
	var updatedBook model.Book
	if err := c.Bind(&updatedBook); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//this is if you want to use ID by URL and not by request body
	//apparently c.Bind only binds the BODY of the payload so the id in the url is not
	//included with the payload
	updatedBook.ID = uint(id)
	book, err := b.BookService.UpdateBook(updatedBook)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, book)
}

func (b *bookController) DeleteBook(c echo.Context) error {
	idParam := c.Param("id")
	id, errID := strconv.ParseUint(idParam, 10, 64)
	if errID != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}
	err := b.BookService.DeleteBook(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Successfully deleted book with ID: %d", id),
	})
}

//func (b *bookController) TestMessage(c echo.Context) error {
//	return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
//}
