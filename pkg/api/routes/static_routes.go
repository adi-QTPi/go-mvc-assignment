package api

import (
	"github.com/adi-QTPi/go-mvc-assignment/pkg/controllers"
	"github.com/gorilla/mux"
)

func ImplimentStaticRoutes(subRouter *mux.Router) {

	renderer := controllers.NewPageRenderer()
	subRouter.HandleFunc("/", renderer.RenderHomePage).Methods("GET")
}
