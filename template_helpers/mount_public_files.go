package template_helpers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var Tmpl *template.Template

func MountPublicFiles(router *mux.Router) {
	fs := http.FileServer(http.Dir("public"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	Tmpl = template.Must(template.New("").Funcs(FuncMap).ParseGlob("./pkg/views/**/*.html"))
}

func MountUploadsFolder(router *mux.Router) {
	fs := http.FileServer(http.Dir("uploads"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/uploads/", fs))
}
