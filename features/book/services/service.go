package services

import (
	"cleanarch/features/book"
	"cleanarch/helper"
	"errors"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

type bookSrvc struct {
	qry book.BookData
	vld *validator.Validate
}

func BookIso(bd book.BookData) book.BookService {
	return &bookSrvc{
		qry: bd,
		vld: validator.New(),
	}
}

func (bs bookSrvc) Add(token interface{}, newBook book.BookCore) (book.BookCore, error) {
	userID := helper.ExtractToken(token)
	if userID == 0 {
		return book.BookCore{}, errors.New("user tidak ditemukan")
	}

	err := bs.vld.Struct(newBook)
	if err != nil {
		if _, v := err.(*validator.ValidationErrors); v {
			log.Println(err)
		}
		return book.BookCore{}, errors.New("validation error")
	}

	res, err := bs.qry.Add(uint(userID), newBook)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "user not found"
		} else {
			msg = "internal server error"
		}
		return book.BookCore{}, errors.New(msg)
	}
	return res, nil
}

func (bs bookSrvc) Update(token interface{}, bookID uint, updateBook book.BookCore) (book.BookCore, error) {
	userID := helper.ExtractToken(token)
	if userID == 0 {
		return book.BookCore{}, errors.New("data tidak ditemukan")
	}

	err := bs.vld.Struct(updateBook)
	if err != nil {
		if _, v := err.(*validator.ValidationErrors); v {
			log.Println(err)
		}
		return book.BookCore{}, errors.New("validation error")
	}

	res, err := bs.qry.Update(uint(userID), bookID, updateBook)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "book not found"
		} else {
			msg = "internal server error"
		}
		return book.BookCore{}, errors.New(msg)
	}

	res.ID = bookID
	res.IDUser = uint(userID)
	return res, nil

}
func (bs bookSrvc) Delete(token interface{}, bookID uint) error {
	userID := helper.ExtractToken(token)
	if userID == 0 {
		return errors.New("user tidak ditemukan")
	}

	err := bs.qry.Delete(uint(userID), bookID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "book not found"
		} else {
			msg = "internal server error"
		}
		return errors.New(msg)
	}
	return nil
}

func (bs bookSrvc) MyBook(token interface{}) ([]book.BookCore, error) {
	userID := helper.ExtractToken(token)

	if userID == 0 {
		return []book.BookCore{}, errors.New("user not found")
	}

	res, err := bs.qry.MyBook(uint(userID))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "book not found"
		} else {
			msg = "internal server error"
		}
		return []book.BookCore{}, errors.New(msg)
	}
	return res, nil
}

func (bs *bookSrvc) BookList() ([]book.BookCore, error) {
	res, err := bs.qry.BookList()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "book not found"
		} else {
			msg = "internal server error"
		}
		return []book.BookCore{}, errors.New(msg)
	}
	return res, nil
}
