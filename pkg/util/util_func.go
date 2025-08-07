package util

import (
	"encoding/json"
	"net/http"
)

func EncodeAndSendResponseWithStatus(w http.ResponseWriter, responseJson StandardResponseJson, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(responseJson)
}
