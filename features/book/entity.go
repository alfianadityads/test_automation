package book

import (
	"github.com/labstack/echo/v4"
)

type BookCore struct {
	ID          uint
	Judul       string `validate:"required"`
	TahunTerbit int    `validate:"required"`
	Penulis     string `validate:"required"`
	IDUser      uint
	Pemilik     string
}

type BookHandler interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	MyBook() echo.HandlerFunc
	BookList() echo.HandlerFunc
}

type BookService interface {
	Add(token interface{}, newBook BookCore) (BookCore, error)
	Update(token interface{}, bookID uint, updateBook BookCore) (BookCore, error)
	Delete(token interface{}, bookID uint) error
	MyBook(token interface{}) ([]BookCore, error)
	BookList() ([]BookCore, error)
}

type BookData interface {
	Add(userID uint, newBook BookCore) (BookCore, error)
	Update(userID uint, bookID uint, updateBook BookCore) (BookCore, error)
	Delete(userID uint, bookID uint) error
	MyBook(userID uint) ([]BookCore, error)
	BookList() ([]BookCore, error)
}
