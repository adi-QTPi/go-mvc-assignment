package controllers

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/config"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
)

type StaticSignupController struct{}

func NewStaticSignupController() *StaticSignupController {
	return &StaticSignupController{}
}

func (sl *StaticSignupController) RenderSignupPage(w http.ResponseWriter, r *http.Request) {
	var responseJson util.StandardResponseJson
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

	err = config.Tmpl.ExecuteTemplate(w, "signup.html", toPage)
	if err != nil {
		responseJson.Msg = "Can't show this page"
		responseJson.ErrDescription = "Error in executing signup.html"
		util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusInternalServerError)
	}
}
