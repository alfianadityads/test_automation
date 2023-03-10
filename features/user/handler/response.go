package handler

import (
	"cleanarch/features/user"
	"net/http"
	"strings"
)

type UserReponse struct {
	ID     uint   `json:"id"`
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	Alamat string `json:"alamat"`
	HP     string `json:"hp"`
}

func ToResponse(data user.UserCore) UserReponse {
	return UserReponse{
		ID:     data.ID,
		Nama:   data.Nama,
		Email:  data.Email,
		Alamat: data.Alamat,
		HP:     data.HP,
	}
}

func PrintSuccessReponse(code int, message string, data ...interface{}) (int, interface{}) {
	resp := map[string]interface{}{}
	if len(data) < 2 {
		resp["data"] = ToResponse(data[0].(user.UserCore))
	} else {
		resp["data"] = ToResponse(data[0].(user.UserCore))
		resp["token"] = data[1].(string)
	}

	if message != "" {
		resp["message"] = message
	}

	return code, resp
}

func PrintErrorResponse(msg string) (int, interface{}) {
	resp := map[string]interface{}{}
	code := -1
	if msg != "" {
		resp["message"] = msg
	}

	if strings.Contains(msg, "server") {
		code = http.StatusInternalServerError
	} else if strings.Contains(msg, "format") {
		code = http.StatusBadRequest
	} else {
		strings.Contains(msg, "not found")
		code = http.StatusNotFound
	}

	return code, resp
}
