package controller

import (
	"TDD-GoAPI/service"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	BookController
}

func New(s *service.Services) *Controller {
	return &Controller{
		BookController: &bookController{s.Book},
	}
}

func SetupRoutes(e *echo.Echo, b BookController) {
	//e.GET("/books", b.TestMessage)
	e.GET("/books", b.GetAllBooks)
	e.GET("/books/:id", b.GetBook)
	e.POST("/books", b.CreateBook)
	e.PUT("/books/:id", b.UpdateBook)
	e.DELETE("/books/:id", b.DeleteBook)
}
