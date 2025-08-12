package controllers

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"github.com/adi-QTPi/go-mvc-assignment/template_helpers"
)

type CookStaticController struct{}

func NewCookStaticController() *CookStaticController {
	return &CookStaticController{}
}

func (cc *CookStaticController) CookDashboardInfo(w http.ResponseWriter, r *http.Request) {

	popup, err := util.ExtractPopupFromFlash(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error extracting popup from flash : %v", err), http.StatusInternalServerError)
	}

	xUser := util.ExtractUserFromContext(r)

	kitchenOrderSlice, err := models.FetchKitchenOrderForToday()
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
	}

	toPage := util.DataToPage{
		Popup:         popup,
		XUser:         xUser,
		KitchenOrders: kitchenOrderSlice,
	}

	var responseJson util.StandardResponseJson
	err = template_helpers.Tmpl.ExecuteTemplate(w, "cook.html", toPage)
	if err != nil {
		responseJson.Msg = "Can't show this page"
		responseJson.ErrDescription = fmt.Sprintf("Error in executing cook.html : %v", err)
		util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusInternalServerError)
	}

	// util.EncodeAndSendKitchenOrderWithStatus(w, kitchenOrderSlice, http.StatusOK)
}
