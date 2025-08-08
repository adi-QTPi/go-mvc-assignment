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

	adminStaticController := controllers.NewAdminStaticController()
	subRouter.Handle("/admin",
		middleware.Chain(
			http.HandlerFunc(adminStaticController.FetchAdminOrderDashboardByDate),
		)).Methods("GET")
	subRouter.Handle("/admin/{date}",
		middleware.Chain(
			http.HandlerFunc(adminStaticController.FetchAdminOrderDashboardByDate),
		)).Methods("GET")

	customerSaticController := controllers.NewCustStaticController()
	subRouter.Handle("/order",
		middleware.Chain(
			http.HandlerFunc(customerSaticController.RenderCustOrderPage),
		)).Methods("GET")
	subRouter.Handle("/order/{date}",
		middleware.Chain(
			http.HandlerFunc(customerSaticController.RenderCustOrderPage),
		)).Methods("GET")
}
