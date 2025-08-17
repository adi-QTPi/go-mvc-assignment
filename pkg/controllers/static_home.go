package controllers

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"github.com/adi-QTPi/go-mvc-assignment/template_helpers"
)

type PageRenderer struct{}

func NewPageRenderer() *PageRenderer {
	return &PageRenderer{}
}

func (pr *PageRenderer) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	xUser := util.ExtractUserFromContext(r)

	toPage := util.DataToPage{
		XUser: xUser,
	}

	err := template_helpers.Tmpl.ExecuteTemplate(w, "homepage.html", toPage)
	if err != nil {
		fmt.Printf("error rendering home page : %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
