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
	var responseJson util.StandardResponseJson
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
		responseJson.Msg = "Can't show this page"
		responseJson.ErrDescription = "Error in executing login.html"
		util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusInternalServerError)
	}
}
