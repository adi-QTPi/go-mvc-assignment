package util

import (
	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/golang-jwt/jwt/v5"
)

type StandardResponseJson struct {
	Msg            string `json:"msg"`
	Err            error  `json:"err"`
	ErrDescription string `json:"err_description"`
}

type JwtCustomClaim struct {
	Sub models.User `json:"sub"`
	jwt.RegisteredClaims
}
