package controllers

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/config"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
)

type StaticController struct{}

func NewStaticController() *StaticController {
	return &StaticController{}
}

func (sl *StaticController) RenderMenuPage(w http.ResponseWriter, r *http.Request) {
	var responseJson util.StandardResponseJson

	xUser := util.ExtractUserFromContext(r)
	categorySlice, err := models.GetAllCategories()
	itemSlice, err := models.GetAllItems()
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

	err = config.Tmpl.ExecuteTemplate(w, "menu.html", toPage)
	if err != nil {
		responseJson.Msg = "Can't show this page"
		responseJson.ErrDescription = fmt.Sprintf("Error in executing menu.html", err)
		util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusInternalServerError)
	}
}
