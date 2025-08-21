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
		fmt.Printf("warning: could not extract popup from flash: %v\n", err)
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
		fmt.Printf("error rendering menu page : %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (sl *StaticController) RenderCartPage(w http.ResponseWriter, r *http.Request) {
	xUser := util.ExtractUserFromContext(r)
	popup, err := util.ExtractPopupFromFlash(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting stuff from sessions and db : %v", err), http.StatusInternalServerError)
		return
	}

	toPage := util.DataToPage{
		Popup: popup,
		XUser: xUser,
	}

	err = template_helpers.Tmpl.ExecuteTemplate(w, "cart.html", toPage)
	if err != nil {
		fmt.Printf("error rendering cart page : %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
