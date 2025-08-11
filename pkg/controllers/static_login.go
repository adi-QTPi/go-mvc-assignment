package controllers

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/config"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
)

type StaticLoginController struct{}

func NewStaticLoginController() *StaticLoginController {
	return &StaticLoginController{}
}

func (sl *StaticLoginController) RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	var responseJson util.StandardResponseJson

	toPage := util.DataToPage{
		Popup: util.ExtractPopupFromContext(r),
		XUser: util.ExtractUserFromContext(r),
	}

	fmt.Print("data to login page", toPage)

	err := config.Tmpl.ExecuteTemplate(w, "login.html", toPage)
	if err != nil {
		responseJson.Msg = "Can't show this page"
		responseJson.ErrDescription = "Error in executing login.html"
		util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusInternalServerError)
	}
}
