package handler

import (
	"cleanarch/features/book"
	"cleanarch/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type bookHdl struct {
	srv book.BookService
}

func BookIso(bs book.BookService) book.BookHandler {
	return &bookHdl{
		srv: bs,
	}
}

func (bh *bookHdl) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddUpdateBookRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		cnv := ReqToCore(input)

		res, err := bh.srv.Add(c.Get("user"), *cnv)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		book := ToResponse("add", res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses menambahkan buku", book))
	}
}

func (bh *bookHdl) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddUpdateBookRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		cnv := ReqToCore(input)

		bookID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		res, err := bh.srv.Update(c.Get("user"), uint(bookID), *cnv)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses mengubah buku", res))
	}
}
func (bh *bookHdl) Delete() echo.HandlerFunc

//func(bh *bookHdl) MyBook() echo.HandlerFunc
