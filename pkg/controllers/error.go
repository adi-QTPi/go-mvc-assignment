package controllers

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/config"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
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

	var responseJson util.StandardResponseJson
	err = config.Tmpl.ExecuteTemplate(w, "error.html", toPage)
	if err != nil {
		responseJson.Msg = "Can't show this page"
		responseJson.ErrDescription = fmt.Sprintf("Error in executing error.html : %v", err)
		util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusInternalServerError)
	}
}
