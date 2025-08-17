package controllers

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"github.com/adi-QTPi/go-mvc-assignment/template_helpers"
)

type StaticSignupController struct{}

func NewStaticSignupController() *StaticSignupController {
	return &StaticSignupController{}
}

func (sl *StaticSignupController) RenderSignupPage(w http.ResponseWriter, r *http.Request) {
	popup, err := util.ExtractPopupFromFlash(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting the popup : %v", err), http.StatusInternalServerError)
		return
	}

	xUser := util.ExtractUserFromContext(r)

	toPage := util.DataToPage{
		Popup: popup,
		XUser: xUser,
	}

	err = template_helpers.Tmpl.ExecuteTemplate(w, "signup.html", toPage)
	if err != nil {
		fmt.Printf("error rendering signup page : %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
