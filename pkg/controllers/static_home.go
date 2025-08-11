package controllers

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/config"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
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

	fmt.Print("data received while rendering home page", toPage)

	err := config.Tmpl.ExecuteTemplate(w, "homepage.html", toPage)
	if err != nil {
		responseJson.Msg = "Can't show this page"
		responseJson.ErrDescription = "Error in executing homepage.html"
		util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusInternalServerError)
	}
}
