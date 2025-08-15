package util

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adi-QTPi/go-mvc-assignment/config"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/golang-jwt/jwt/v5"
)

func GetSignedJwtOfUser(w http.ResponseWriter, userId string) (string, error) {
	config.LoadJwtEnv()

	user, err := models.GetUserById(userId)
	if err != nil {
		return "", err
	}

	claims := JwtCustomClaim{
		Sub: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedTokenString, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return signedTokenString, nil
}

func DecryptJwtToken(w http.ResponseWriter, r *http.Request, tokenValue string) models.User {
	config.LoadJwtEnv()

	claims := &JwtCustomClaim{}

	token, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JWT_SECRET), nil
	})

	if err != nil || !token.Valid {
		var blankUser models.User
		popup := Popup{
			Msg:     "Invalid or Expired token, login again",
			IsError: true,
		}
		InsertPopupInFlash(w, r, popup)

		RedirectToSite(w, r, "/login")
		return blankUser
	}

	return claims.Sub
}
