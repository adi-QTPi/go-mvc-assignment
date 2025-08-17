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

	err = template_helpers.Tmpl.ExecuteTemplate(w, "cook.html", toPage)
	if err != nil {
		fmt.Printf("error rendering cook page : %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
