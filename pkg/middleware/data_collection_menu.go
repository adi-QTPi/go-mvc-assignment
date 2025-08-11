package middleware

import (
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
)

func FetchMenu(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		util.RedirectToSite(w, r, "/api/items")
	})
}

func FetchCategories(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		util.RedirectToSite(w, r, "/api/categories")
	})
}
