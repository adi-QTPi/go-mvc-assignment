package api

import (
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/controllers"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/middleware"
	"github.com/gorilla/mux"
)

func ImplimentStaticRoutes(subRouter *mux.Router) {

	subRouter.Use(middleware.IdentifyUser)

	renderer := controllers.NewPageRenderer()
	subRouter.HandleFunc("/", renderer.RenderHomePage).Methods("GET")

	cookController := controllers.NewCookStaticController()
	subRouter.Handle("/cook",
		middleware.Chain(
			http.HandlerFunc(cookController.CookDashboardInfo),
		)).Methods("GET")
}
