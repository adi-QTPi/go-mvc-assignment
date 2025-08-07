package util

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/adi-QTPi/go-mvc-assignment/config"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/golang-jwt/jwt/v5"
)

func EncodeAndSendResponseWithStatus(w http.ResponseWriter, responseJson StandardResponseJson, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(responseJson)
}

func PutInContext(r *http.Request, key string, value string) *http.Request {
	ctx := context.WithValue(r.Context(), key, value)
	r = r.WithContext(ctx)

	return r
}

func PutUserInContext(r *http.Request, user models.User) *http.Request {
	ctx := context.WithValue(r.Context(), "xUser", user)
	return r.WithContext(ctx)
}

func ExtractFromContext(r *http.Request, key string) string {
	val, _ := r.Context().Value(key).(string)
	return val
}

func ExtractUserFromContext(r *http.Request) models.User {
	val, _ := r.Context().Value("xUser").(models.User)
	return val
}

func GetSignedJwtOfUser(w http.ResponseWriter, userId string) (string, error) {
	config.LoadJwtEnv()

	user, err := models.GetUserById(userId)
	if err != nil {
		return "", err
	}

	userDereference := *user

	claims := JwtCustomClaim{
		Sub: userDereference,
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
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		var blankUser models.User
		return blankUser
	}

	return claims.Sub
}
