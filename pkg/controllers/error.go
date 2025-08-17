package controllers

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"github.com/adi-QTPi/go-mvc-assignment/template_helpers"
)

type StaticErrorController struct{}

func NewStaticErrorController() *StaticErrorController {
	return &StaticErrorController{}
}

func (asc *StaticErrorController) RenderErrorPage(w http.ResponseWriter, r *http.Request) {
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

	err = template_helpers.Tmpl.ExecuteTemplate(w, "error.html", toPage)
	if err != nil {
		fmt.Printf("error rendering error page : %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
