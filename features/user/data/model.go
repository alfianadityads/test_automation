package data

import (
	"cleanarch/features/user"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Nama     string
	Email    string
	Alamat   string
	HP       string
	Password string
}

func ModelToCore(data UserModel) user.UserCore {
	return user.UserCore{
		ID:       data.ID,
		Nama:     data.Nama,
		Email:    data.Email,
		Alamat:   data.Alamat,
		HP:       data.HP,
		Password: data.Password,
	}
}

func CoreToModel(data user.UserCore) UserModel {
	return UserModel{
		Model:    gorm.Model{ID: data.ID},
		Nama:     data.Nama,
		Email:    data.Email,
		Alamat:   data.Alamat,
		HP:       data.HP,
		Password: data.Password,
	}

}
