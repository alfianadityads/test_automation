package data

import (
	"cleanarch/features/user"
	"errors"
	"log"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func Isolation(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

// fungsi register pada user data menerima param input newUser dan return ke user core, error
func (uq *userQuery) Register(newUser user.UserCore) (user.UserCore, error) {
	cnv := CoreToModel(newUser)
	err := uq.db.Create(&cnv).Error
	if err != nil {
		return user.UserCore{}, err
	}

	newUser.ID = cnv.ID

	return newUser, nil
}

// fungsi login menerima input email dan mengembalikan user core, error
func (uq *userQuery) Login(email string) (user.UserCore, error) {
	res := UserModel{}

	err := uq.db.Where("email= ?", email).First(&res).Error
	if err != nil {
		log.Println("login query error", err.Error())
		return user.UserCore{}, errors.New("user data not found")
	}

	return ModelToCore(res), nil
}

// fungsi show profile menerima input id dan mengembalikan user core, error
func (uq *userQuery) Profile(id uint) (user.UserCore, error) {
	res := UserModel{}
	err := uq.db.Where("id = ?", id).First(&res).Error
	if err != nil {
		log.Println("get by ID query error", err.Error())
		return user.UserCore{}, errors.New("user data not found")
	}

	return ModelToCore(res), nil
}

// 
func (uq *userQuery) Update(userID uint, updateUser user.UserCore) (user.UserCore, error) {
	cnv := CoreToModel(updateUser)
	qry := uq.db.Model(&UserModel{}).Where("id = ?", userID).Updates(&cnv)

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return user.UserCore{}, errors.New("tidak ada data user yang diubah")
	}

	err := qry.Error
	if err != nil {
		log.Println("update query error")
		return user.UserCore{}, errors.New("tidak bisa mengubah data user")
	}

	return ModelToCore(cnv), nil
}

func (uq *userQuery) Deactive(userID uint) error {
	qry := uq.db.Where("id = ?", userID).Delete(&UserModel{})

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return errors.New("tidak ada data user yang terhapus")
	}

	err := qry.Error
	if err != nil {
		log.Println("delete quey error")
		return errors.New("tidak bisa menghapus data user")
	}

	return nil
}
