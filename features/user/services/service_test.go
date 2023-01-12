package services

// import (
// 	"cleanarch/features/user"
// 	"cleanarch/helper"
// 	"cleanarch/mocks"
// 	"errors"
// 	"testing"

// 	"github.com/golang-jwt/jwt"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestLogin(t *testing.T) {
// 	repo := mocks.NewUserData(t) // mock data

// 	t.Run("Berhasil login", func(t *testing.T) {
// 		// input dan respond untuk mock data
// 		inputEmail := "jerry@alterra.id"
// 		// res dari data akan mengembalikan password yang sudah di hash
// 		hashed, _ := helper.GeneratePassword("be1422")
// 		resData := user.UserCore{ID: uint(1), Nama: "jerry", Email: "jerry@alterra.id", HP: "08123456", Password: hashed}

// 		repo.On("Login", inputEmail).Return(resData, nil) // simulasi method login pada layer data

// 		srv := Isolation(repo)
// 		token, res, err := srv.Login(inputEmail, "be1422")
// 		assert.Nil(t, err)
// 		assert.NotEmpty(t, token)
// 		assert.Equal(t, resData.ID, res.ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Tidak ditemukan", func(t *testing.T) {
// 		inputEmail := "putra@alterra.id"
// 		repo.On("Login", inputEmail).Return(user.UserCore{}, errors.New("data not found"))

// 		srv := Isolation(repo)
// 		token, res, err := srv.Login(inputEmail, "be1422")
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "tidak ditemukan")
// 		assert.Empty(t, token)
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Salah password", func(t *testing.T) {
// 		inputEmail := "jerry@alterra.id"
// 		hashed, _ := helper.GeneratePassword("be1422")
// 		resData := user.UserCore{ID: uint(1), Nama: "jerry", Email: "jerry@alterra.id", HP: "08123456", Password: hashed}
// 		repo.On("Login", inputEmail).Return(resData, nil)

// 		srv := Isolation(repo)
// 		token, res, err := srv.Login(inputEmail, "be1423")
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "password tidak sesuai")
// 		assert.Empty(t, token)
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})

// }

// func TestProfile(t *testing.T) {
// 	repo := mocks.NewUserData(t)

// 	t.Run("Sukses lihat profile", func(t *testing.T) {
// 		resData := user.UserCore{ID: uint(1), Nama: "jerry", Email: "jerry@alterra.id", HP: "08123456"}

// 		repo.On("Profile", uint(1)).Return(resData, nil).Once()

// 		srv := Isolation(repo)

// 		_, token := helper.GenerateJWT(1)

// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true

// 		res, err := srv.Profile(pToken)
// 		assert.Nil(t, err)
// 		assert.Equal(t, resData.ID, res.ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("jwt tidak valid", func(t *testing.T) {
// 		srv := Isolation(repo)

// 		_, token := helper.GenerateJWT(1)

// 		res, err := srv.Profile(token)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "not found")
// 		assert.Equal(t, uint(0), res.ID)
// 	})

// 	t.Run("data tidak ditemukan", func(t *testing.T) {
// 		repo.On("Profile", uint(4)).Return(user.UserCore{}, errors.New("data not found")).Once()

// 		srv := Isolation(repo)

// 		_, token := helper.GenerateJWT(4)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Profile(pToken)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "tidak ditemukan")
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("masalah di server", func(t *testing.T) {
// 		repo.On("Profile", mock.Anything).Return(user.UserCore{}, errors.New("terdapat masalah pada server")).Once()
// 		srv := Isolation(repo)

// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Profile(pToken)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "server")
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})
// }

// // func TestRegister(t *testing.T) {
// // 	repo := mocks.NewUserData(t)
// // }

// func TestUpdate(t *testing.T) {
// 	repo := mocks.NewUserData(t)

// 	t.Run("sukses update data", func(t *testing.T) {
// 		inputData := user.UserCore{Nama: "alfian", Email: "alfian@example.com", HP: "081234567"}

// 		hash, _ := helper.GeneratePassword("alfian1221")
// 		resData := user.UserCore{ID: uint(1), Nama: "alfian", Email: "alfian@example.com", HP: "081234567", Password: hash}
// 		repo.On("Update", uint(1), inputData).Return(resData, nil)

// 		srv := Isolation(repo)

// 		_, token := helper.GenerateJWT(1)

// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true

// 		res, err := srv.Update(pToken, inputData)
// 		assert.Nil(t, err)
// 		assert.Equal(t, resData.ID, res.ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("tidak ditemukan", func(t *testing.T) {
// 		inputData := user.UserCore{Nama: "alfian", Email: "alfian@example.com", HP: "081234567"}
// 		repo.On("Update", uint(2), inputData).Return(user.UserCore{}, errors.New("not found")).Once()

// 		srv := Isolation(repo)

// 		_, token := helper.GenerateJWT(2)

// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true

// 		res, err := srv.Update(pToken, inputData)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "tidak ditemukan")
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("masalah pada server", func(t *testing.T) {
// 		inputData := user.UserCore{Nama: "alfian", Email: "alfian@example.com", HP: "081234567"}
// 		repo.On("Update", uint(2), inputData).Return(user.UserCore{}, errors.New("terdapat masalah pada server")).Once()

// 		srv := Isolation(repo)

// 		_, token := helper.GenerateJWT(2)

// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true

// 		res, err := srv.Update(pToken, inputData)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "server")
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("jwt tidak valid", func(t *testing.T) {
// 		inputData := user.UserCore{Nama: "alfian", Email: "alfian@example.com", HP: "081234567"}

// 		srv := Isolation(repo)

// 		_, token := helper.GenerateJWT(0)

// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true

// 		res, err := srv.Update(pToken, inputData)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "data not found")
// 		assert.Equal(t, uint(0), res.ID)
// 	})
// }
