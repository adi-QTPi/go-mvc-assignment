package util

import (
	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/golang-jwt/jwt/v5"
)

type StandardResponseJson struct {
	Msg            string `json:"msg"`
	ErrDescription string `json:"err_description"`
}

type JwtCustomClaim struct {
	Sub models.User `json:"sub"`
	jwt.RegisteredClaims
}

type Popup struct {
	Msg     string
	IsError bool
}

type DataToPage struct {
	XUser models.User
	Popup Popup
}
