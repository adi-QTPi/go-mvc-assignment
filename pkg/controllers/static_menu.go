package controllers

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"github.com/adi-QTPi/go-mvc-assignment/template_helpers"
)

type StaticController struct{}

func NewStaticController() *StaticController {
	return &StaticController{}
}

func (sl *StaticController) RenderMenuPage(w http.ResponseWriter, r *http.Request) {
	var responseJson util.StandardResponseJson

	xUser := util.ExtractUserFromContext(r)
	categorySlice, err := models.GetAllCategories()
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting category slice from sessions and db : %v", err), http.StatusInternalServerError)
		return
	}
	itemSlice, err := models.GetAllItems()
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting itemSlice from sessions and db : %v", err), http.StatusInternalServerError)
		return
	}
	popup, err := util.ExtractPopupFromFlash(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting stuff from sessions and db : %v", err), http.StatusInternalServerError)
		return
	}

	toPage := util.DataToPage{
		Popup:         popup,
		XUser:         xUser,
		CategorySlice: categorySlice,
		ItemSlice:     itemSlice,
	}

	err = template_helpers.Tmpl.ExecuteTemplate(w, "menu.html", toPage)
	if err != nil {
		responseJson.Msg = "Can't show this page"
		responseJson.ErrDescription = fmt.Sprintf("Error in executing menu.html : %v", err)
		util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusInternalServerError)
	}
}

func (sl *StaticController) RenderCartPage(w http.ResponseWriter, r *http.Request) {
	var responseJson util.StandardResponseJson

	xUser := util.ExtractUserFromContext(r)
	// categorySlice, err := models.GetAllCategories()
	// itemSlice, err := models.GetAllItems()
	popup, err := util.ExtractPopupFromFlash(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting stuff from sessions and db : %v", err), http.StatusInternalServerError)
		return
	}

	toPage := util.DataToPage{
		Popup: popup,
		XUser: xUser,
		// CategorySlice: categorySlice,
		// ItemSlice:     itemSlice,
	}

	err = template_helpers.Tmpl.ExecuteTemplate(w, "cart.html", toPage)
	if err != nil {
		responseJson.Msg = "Can't show this page"
		responseJson.ErrDescription = fmt.Sprintf("Error in executing menu.html : %v", err)
		util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusInternalServerError)
	}
}
