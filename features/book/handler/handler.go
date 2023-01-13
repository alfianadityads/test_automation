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
func (bh *bookHdl) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		paramID := c.Param("id")
		bookID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("Convert id error : ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Please input number only",
			})
		}

		err = bh.srv.Delete(token, uint(bookID))

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusAccepted, "deleted a book successfully")
	}
}

func (bh *bookHdl) MyBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := bh.srv.BookList()
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "displayed all books successfully", res))
	}
}

func (bh *bookHdl) BookList() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := bh.srv.BookList()
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "display all books completed", res))
	}
}
