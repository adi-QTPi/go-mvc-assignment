package util

import (
	"context"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
)

func PutInContext(r *http.Request, key string, value string) *http.Request {
	ctx := context.WithValue(r.Context(), key, value)
	r = r.WithContext(ctx)

	return r
}

func PutUserInContext(r *http.Request, user models.User) *http.Request {
	ctx := context.WithValue(r.Context(), "xUser", user)
	return r.WithContext(ctx)
}

func PutCartInContext(r *http.Request, orderSlice []models.ItemInCart) *http.Request {
	ctx := context.WithValue(r.Context(), "orderSlice", orderSlice)
	return r.WithContext(ctx)
}

func PutPopupInContext(r *http.Request, popup Popup) *http.Request {
	ctx := context.WithValue(r.Context(), "popup", popup)
	return r.WithContext(ctx)
}

func ExtractPopupFromContext(r *http.Request) Popup {
	val, _ := r.Context().Value("popup").(Popup)
	return val
}

func ExtractFromContext(r *http.Request, key string) string {
	val, _ := r.Context().Value(key).(string)
	return val
}

func ExtractUserFromContext(r *http.Request) models.User {
	val, _ := r.Context().Value("xUser").(models.User)
	return val
}
func ExtractCartFromContext(r *http.Request) []models.ItemInCart {
	val, _ := r.Context().Value("orderSlice").([]models.ItemInCart)
	return val
}
