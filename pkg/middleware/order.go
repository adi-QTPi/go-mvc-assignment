package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
)

func AssignEmptyTable(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xUser := util.ExtractUserFromContext(r)

		tableNo, err := models.CheckAndAssignTable(xUser.UserId)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching categories: %v", err), http.StatusInternalServerError)
			return
		}
		if tableNo == 0 {
			var responseJson util.StandardResponseJson
			responseJson.ErrDescription = "No empty tables available at this time"
			util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusOK)
			return
		}

		tableNoStr := strconv.FormatInt(tableNo, 10)

		r = util.PutInContext(r, "tableNo", tableNoStr)
		next.ServeHTTP(w, r)
	})
}

func DecodeCartJsonInput(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var orderSlice []models.ItemInCart

		err := json.NewDecoder(r.Body).Decode(&orderSlice)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid farmatting, from Decode Cart JSON Input middleware : %v", err), http.StatusBadRequest)
			return
		}

		r = util.PutCartInContext(r, orderSlice)

		next.ServeHTTP(w, r)
	})
}
