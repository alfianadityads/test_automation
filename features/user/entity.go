package user

import (
	"github.com/labstack/echo/v4"
)

type UserCore struct {
	ID       uint
	Nama     string
	Email    string
	Alamat   string
	HP       string
	Password string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Deactive() echo.HandlerFunc
}

type UserService interface {
	Register(newUser UserCore) (UserCore, error)
	Login(email, password string) (string, UserCore, error)
	Profile(token interface{}) (UserCore, error)
	Update(token interface{}, updateUser UserCore) (UserCore, error)
	Deactive(token interface{}) error
}

type UserData interface {
	Register(newUser UserCore) (UserCore, error)
	Login(email string) (UserCore, error)
	Profile(id uint) (UserCore, error)
	Update(id uint, updateUser UserCore) (UserCore, error)
	Deactive(id uint) error
}
