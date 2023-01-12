package data

import (
	"cleanarch/features/book"

	"gorm.io/gorm"
)

type BookModel struct {
	gorm.Model
	Judul       string
	TahunTerbit int
	Penulis     string
	IDUser      uint
}

type BookPemilik struct {
	ID          uint
	Judul       string
	TahunTerbit int
	Penulis     string
	Pemilik     string
}

func BookToCore(data BookModel) book.BookCore {
	return book.BookCore{
		ID:          data.ID,
		Judul:       data.Judul,
		TahunTerbit: data.TahunTerbit,
		Penulis:     data.Penulis,
		IDUser:      data.IDUser,
	}
}

func CoreToBook(data book.BookCore) BookModel {
	return BookModel{
		Model:       gorm.Model{ID: data.ID},
		Judul:       data.Judul,
		Penulis:     data.Penulis,
		TahunTerbit: data.TahunTerbit,
		IDUser:      data.IDUser,
	}
}

func (dataModel *BookModel) ModelToCore() book.BookCore {
	return book.BookCore{
		ID:          dataModel.ID,
		Judul:       dataModel.Judul,
		Penulis:     dataModel.Penulis,
		TahunTerbit: dataModel.TahunTerbit,
		IDUser:      dataModel.IDUser,
	}
}

func ListModelToCore(dataModel []BookModel) []book.BookCore {
	var dataCore []book.BookCore
	for _, val := range dataModel {
		dataCore = append(dataCore, val.ModelToCore())
	}
	return dataCore
}

func (dataModel *BookPemilik) BookPemilikToCore() book.BookCore {
	return book.BookCore{
		ID:          dataModel.ID,
		Judul:       dataModel.Judul,
		TahunTerbit: dataModel.TahunTerbit,
		Penulis:     dataModel.Penulis,
		Pemilik:     dataModel.Pemilik,
	}
}

func ListPemilikToCore(dataModel []BookPemilik) []book.BookCore {
	var dataCore []book.BookCore
	for _, val := range dataModel {
		dataCore = append(dataCore, val.BookPemilikToCore())
	}
	return dataCore
}
