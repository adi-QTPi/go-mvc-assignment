package config

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var Tmpl *template.Template

func MountPublicFiles(router *mux.Router) {
	fs := http.FileServer(http.Dir("public"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	Tmpl = template.Must(template.New("").Funcs(FuncMap).ParseGlob("pkg/views/**/*.html"))
}

var FuncMap = template.FuncMap{
	"ToJSON": ToJSON,
}

func ToJSON(v interface{}) (template.JS, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return template.JS(b), nil
}
