package api

import (
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/controllers"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/middleware"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"github.com/gorilla/mux"
)

func ImplementRootRoutes(subRouter *mux.Router) {
	staticLoginController := controllers.NewStaticLoginController()
	subRouter.Handle("/login",
		middleware.Chain(
			http.HandlerFunc(staticLoginController.RenderLoginPage),
		)).Methods("GET")

	subRouter.Handle("/",
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				util.RedirectToSite(w, r, "/static")
			},
		))

}
