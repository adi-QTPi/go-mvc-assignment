package controllers

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"github.com/adi-QTPi/go-mvc-assignment/template_helpers"
)

type StaticLoginController struct{}

func NewStaticLoginController() *StaticLoginController {
	return &StaticLoginController{}
}

func (sl *StaticLoginController) RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	popup, err := util.ExtractPopupFromFlash(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting the token : %v", err), http.StatusInternalServerError)
		return
	}

	toPage := util.DataToPage{
		Popup: popup,
	}

	err = template_helpers.Tmpl.ExecuteTemplate(w, "login.html", toPage)
	if err != nil {
		fmt.Printf("error rendering login page : %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
