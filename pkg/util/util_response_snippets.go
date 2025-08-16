package util

import (
	"encoding/json"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
)

func EncodeAndSendResponseWithStatus(w http.ResponseWriter, responseJson StandardResponseJson, statusCode int) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
	}
	json.NewEncoder(w).Encode(responseJson)
}

func EncodeAndSendUsersWithStatus(w http.ResponseWriter, statusCode int, users ...models.User) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(users)
}
func EncodeAndSendCategoriesWithStatus(w http.ResponseWriter, catSlice []models.Category, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(catSlice)
}

func EncodeAndSendItemWithStatus(w http.ResponseWriter, itemSlice []models.DisplayItem, statusCode int) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
	}
	json.NewEncoder(w).Encode(itemSlice)
}

func EncodeAndSendKitchenOrderWithStatus(w http.ResponseWriter, itemSlice []models.KitchenOrder, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(itemSlice)
}

func EncodeAndSendOrderWithStatus(w http.ResponseWriter, itemSlice []models.Order, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(itemSlice)
}
