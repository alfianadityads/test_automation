package data

import (
	"cleanarch/features/book"
	"errors"
	"log"

	"gorm.io/gorm"
)

type bookQuery struct {
	db *gorm.DB
}

func BookIso(db *gorm.DB) book.BookData {
	return &bookQuery{
		db: db,
	}
}

func (bq *bookQuery) Add(userID uint, newBook book.BookCore) (book.BookCore, error) {
	cnv := CoreToBook(newBook)
	cnv.IDUser = uint(userID)
	err := bq.db.Create(&cnv).Error
	if err != nil {
		return book.BookCore{}, err
	}
	newBook.ID = cnv.ID

	return newBook, nil
}

func (bq *bookQuery) Update(UserID uint, bookID uint, updateBook book.BookCore) (book.BookCore, error) {
	getID := BookModel{}
	err1 := bq.db.Where("id = ?", bookID).First(&getID).Error

	if err1 != nil {
		log.Println("get book error : ", err1.Error())
		return book.BookCore{}, err1
	}

	cnv := CoreToBook(updateBook)
	qry := bq.db.Where("id = ?", bookID).Updates(&cnv)

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return book.BookCore{}, errors.New("tidak ada data buku yang diubah")
	}

	err := qry.Error
	if err != nil {
		log.Println("update quey error")
		return book.BookCore{}, errors.New("tidak bisa merubah data buku")
	}
	return BookToCore(cnv), nil
}

func (bq *bookQuery) Delete(userID uint, bookID uint) error {
	qry := bq.db.Where("id = ? AND id_user = ?", bookID, userID).Delete(&BookModel{})

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return errors.New("tidak ada buku yang dihapus")
	}

	err := qry.Error
	if err != nil {
		log.Println("delete quey error")
		return errors.New("tidak bisa menghapus data buku")
	}
	return nil
}

func (bq *bookQuery) MyBook(userID uint) ([]book.BookCore, error) {
	res := []BookModel{}
	if err := bq.db.Where("user_id = ?", userID).Find(&res).Error; err != nil {
		log.Println("get book by ID query error : ", err.Error())
		return []book.BookCore{}, err
	}
	return ListModelToCore(res), nil
}

func (bq *bookQuery) BookList() ([]book.BookCore, error) {
	res := []BookPemilik{}
	if err := bq.db.Table("books").Joins("JOIN users ON users.id = books.user_id").Select("books.id, books.title, books.year, books.author, users.name AS owner").Find(&res).Error; err != nil {
		log.Println("Get all books query error : ", err.Error())
		return []book.BookCore{}, err
	}

	return ListPemilikToCore(res), nil
}
