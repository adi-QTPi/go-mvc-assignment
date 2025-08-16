package controllers

import (
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"github.com/adi-QTPi/go-mvc-assignment/template_helpers"
)

type PageRenderer struct{}

func NewPageRenderer() *PageRenderer {
	return &PageRenderer{}
}

func (pr *PageRenderer) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	var responseJson util.StandardResponseJson
	xUser := util.ExtractUserFromContext(r)

	toPage := util.DataToPage{
		XUser: xUser,
	}

	err := template_helpers.Tmpl.ExecuteTemplate(w, "homepage.html", toPage)
	if err != nil {
		responseJson.Msg = "Can't show this page"
		responseJson.ErrDescription = "Error in executing homepage.html"
		util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusInternalServerError)
		return
	}
}
