package handler

import (
	"cleanarch/features/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControll struct {
	srv user.UserService
}

func Isolation(srv user.UserService) user.UserHandler {
	return &UserControll{
		srv: srv,
	}
}

func (uc *UserControll) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format input salah")
		}

		res, err := uc.srv.Register(*ReqToCore(input))
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessReponse(http.StatusCreated, "berhasil mendaftarkan akun", res))
	}
}

func (uc *UserControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format input tidak sesuai")
		}

		token, res, err := uc.srv.Login(input.Email, input.Password)
		if err != nil {
			return c.JSON((PrintErrorResponse(err.Error())))
		}
		return c.JSON(PrintSuccessReponse(http.StatusOK, "berhasil login", res, token))
	}
}

func (uc *UserControll) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := uc.srv.Profile(token)
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessReponse(http.StatusOK, "berhasil menampilkan profile", res))
	}
}

func (uc *UserControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := UpdateRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format input tidak sesuai")
		}

		res, err := uc.srv.Update(token, *ReqToCore(input))
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessReponse(http.StatusOK, "berhasil update data profile", res))
	}
}

func (uc *UserControll) Deactive() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		err := uc.srv.Deactive(token)
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusAccepted, "berhasil menghapus data user")
	}
}
