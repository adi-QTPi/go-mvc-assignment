package config

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var Tmpl *template.Template

func MountPublicFiles(router *mux.Router) {
	fs := http.FileServer(http.Dir("public"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	Tmpl = template.Must(template.ParseGlob("pkg/views/**/*.html"))
}
