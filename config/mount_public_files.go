package config

import (
	"html/template"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/template_helpers"
	"github.com/gorilla/mux"
)

var Tmpl *template.Template

func MountPublicFiles(router *mux.Router) {
	fs := http.FileServer(http.Dir("public"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	Tmpl = template.Must(template.New("").Funcs(FuncMap).ParseGlob("pkg/views/**/*.html"))
}

var FuncMap = template.FuncMap{
	"ToJSON":         template_helpers.ToJSON,
	"add":            template_helpers.Add,
	"multiply":       template_helpers.Multiply,
	"booleanUpdater": template_helpers.CookPageHelper,
}
