package controllers

import (
	"fmt"
	"net/http"
)

type PageRenderer struct{}

func NewPageRenderer() *PageRenderer {
	return &PageRenderer{}
}

func (pr *PageRenderer) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "showing the home page...")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error Showing page: %v", err), http.StatusInternalServerError)
		return
	}
}
