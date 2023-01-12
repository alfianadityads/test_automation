package handler

import (
	"cleanarch/features/book"
)

type AddUpdateBookRequest struct {
	Judul       string `json:"judul"`
	TahunTerbit int    `json:"tahun_terbit"`
	Penulis     string `json:"penulis"`
}

func ReqToCore(data interface{}) *book.BookCore {
	res := book.BookCore{}

	switch data.(type) {
	case AddUpdateBookRequest:
		cnv := data.(AddUpdateBookRequest)
		res.Judul = cnv.Judul
		res.TahunTerbit = cnv.TahunTerbit
		res.Penulis = cnv.Penulis
	default:
		return nil
	}

	return &res
}