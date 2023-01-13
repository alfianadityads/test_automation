package services

import (
	"cleanarch/features/book"
	"cleanarch/helper"
	"cleanarch/mocks"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewBookData(t)

	t.Run("sukses menambahkan buku", func(t *testing.T) {
		input := book.BookCore{Judul: "kariage-kun", TahunTerbit: 1998, Penulis: "Masashi Ueda"}
		resData := book.BookCore{ID: uint(1), Judul: "kariage-kun", TahunTerbit: 1998, Penulis: "Masashi Ueda"}
		repo.On("Add", uint(1), input).Return(resData, nil).Once()

		srv := BookIso(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, input)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		// repo.AssertExpectations(t)
	})

	t.Run("user tidak ditemukan", func(t *testing.T) {
		input := book.BookCore{Judul: "kariage-kun", TahunTerbit: 1998, Penulis: "Masashi Ueda"}
		repo.On("Add", uint(1), input).Return(book.BookCore{}, errors.New("data not found")).Once()

		srv := BookIso(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, input)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		input := book.BookCore{Judul: "kariage-kun", TahunTerbit: 1998, Penulis: "Masashi Ueda"}
		repo.On("Add", uint(1), input).Return(book.BookCore{}, errors.New("terdapat masalah pada server")).Once()

		srv := BookIso(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, input)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		input := book.BookCore{Judul: "kariage-kun", TahunTerbit: 1998, Penulis: "Masashi Ueda"}

		srv := BookIso(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, input)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewBookData(t)

	t.Run("sukses update data", func(t *testing.T) {
		input := book.BookCore{Judul: "kariage-kun", TahunTerbit: 1998, Penulis: "Masashi Ueda"}
		resBook := book.BookCore{ID: uint(1), Judul: "kariage-kun", TahunTerbit: 1998, Penulis: "Masashi Ueda"}
		repo.On("Update", uint(1), uint(1), input).Return(resBook, nil).Once()

		srv := BookIso(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, uint(1), input)
		assert.Nil(t, err)
		assert.Equal(t, resBook.ID, res.ID)
		assert.Equal(t, input.Judul, res.Judul)
		assert.Equal(t, input.TahunTerbit, res.TahunTerbit)
		assert.Equal(t, input.Penulis, res.Penulis)
		repo.AssertExpectations(t)

	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		input := book.BookCore{Judul: "kariage-kun", TahunTerbit: 1998, Penulis: "Masashi Ueda"}

		srv := BookIso(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, 1, input)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		input := book.BookCore{Judul: "kariage-kun", TahunTerbit: 1998, Penulis: "Masashi Ueda"}
		repo.On("Update", uint(5), uint(5), input).Return(book.BookCore{}, errors.New("data not found")).Once()

		srv := BookIso(repo)
		_, token := helper.GenerateJWT(5)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, 5, input)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		input := book.BookCore{Judul: "kariage-kun", TahunTerbit: 1998, Penulis: "Masashi Ueda"}
		repo.On("Update", uint(1), uint(1), input).Return(book.BookCore{}, errors.New("terdapat masalah pada server")).Once()

		srv := BookIso(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, 1, input)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewBookData(t)

	t.Run("sukses menghapus book", func(t *testing.T) {
		repo.On("Delete", uint(1), uint(1)).Return(nil).Once()

		srv := BookIso(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := BookIso(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(token, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		repo.AssertExpectations(t)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		repo.On("Delete", uint(5), uint(1)).Return(errors.New("data not found")).Once()

		srv := BookIso(repo)

		_, token := helper.GenerateJWT(5)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		repo.On("Delete", uint(1), uint(1)).Return(errors.New("terdapat masalah pada server")).Once()
		srv := BookIso(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

}

func TestMyBook(t *testing.T) {
	repo := mocks.NewBookData(t)

	t.Run("Sukses lihat buku", func(t *testing.T) {
		resData := []book.BookCore{
			{
				Judul:       "Kariage-kun",
				TahunTerbit: 1998,
				Penulis:     "Masashi Ueda",
				Pemilik:     "Alfian",
			},
		}

		repo.On("MyBook", uint(1)).Return(resData, nil).Once()

		srv := BookIso(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.MyBook(pToken)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := BookIso(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		_, err := srv.MyBook(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		repo.On("MyBook", uint(5)).Return([]book.BookCore{}, errors.New("data not found")).Once()

		srv := BookIso(repo)

		_, token := helper.GenerateJWT(5)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.MyBook(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, res, []book.BookCore{})
		repo.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		repo.On("MyBook", uint(1)).Return([]book.BookCore{}, errors.New("terdapat masalah pada server")).Once()
		srv := BookIso(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.MyBook(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res, []book.BookCore{})
		repo.AssertExpectations(t)
	})
}

func TestBookList(t *testing.T) {
	repo := mocks.NewBookData(t)

	t.Run("Sukses lihat buku", func(t *testing.T) {
		resData := []book.BookCore{
			{
				Judul:       "Kariage-kun",
				TahunTerbit: 1998,
				Penulis:     "Masashi Ueda",
				Pemilik:     "Alfian",
			},
		}

		repo.On("BookList").Return(resData, nil).Once()

		srv := BookIso(repo)

		res, err := srv.BookList()
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		repo.On("BookList").Return([]book.BookCore{}, errors.New("data not found")).Once()

		srv := BookIso(repo)
		res, err := srv.BookList()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		repo.On("BookList").Return([]book.BookCore{}, errors.New("terdapat masalah pada server")).Once()
		srv := BookIso(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.BookList()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}
