package controllers

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
)

type CookStaticController struct{}

func NewCookStaticController() *CookStaticController {
	return &CookStaticController{}
}

func (cc *CookStaticController) CookDashboardInfo(w http.ResponseWriter, r *http.Request) {

	kitchenOrderSlice, err := models.FetchKitchenOrderForToday()
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
	}

	util.EncodeAndSendKitchenOrderWithStatus(w, kitchenOrderSlice, http.StatusOK)
}
