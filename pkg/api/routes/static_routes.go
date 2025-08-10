package api

import (
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/controllers"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/middleware"
	"github.com/gorilla/mux"
)

func ImplementStaticRoutes(subRouter *mux.Router) {

	subRouter.Use(middleware.IdentifyUser)

	renderer := controllers.NewPageRenderer()
	subRouter.HandleFunc("/", renderer.RenderHomePage).Methods("GET")

	cookController := controllers.NewCookStaticController()
	subRouter.Handle("/cook",
		middleware.Chain(
			http.HandlerFunc(cookController.CookDashboardInfo),
			middleware.RestrictToRoles("cook"),
		)).Methods("GET")

	adminStaticController := controllers.NewAdminStaticController()
	subRouter.Handle("/admin",
		middleware.Chain(
			http.HandlerFunc(adminStaticController.FetchAdminOrderDashboardByDate),
			middleware.RestrictToRoles("admin"),
		)).Methods("GET")
	subRouter.Handle("/admin/{date}",
		middleware.Chain(
			http.HandlerFunc(adminStaticController.FetchAdminOrderDashboardByDate),
			middleware.RestrictToRoles("admin"),
		)).Methods("GET")

	customerSaticController := controllers.NewCustStaticController()
	subRouter.Handle("/order",
		middleware.Chain(
			http.HandlerFunc(customerSaticController.RenderCustOrderPage),
			middleware.RestrictToRoles("customer"),
		)).Methods("GET")
	subRouter.Handle("/order/{date}",
		middleware.Chain(
			http.HandlerFunc(customerSaticController.RenderCustOrderPage),
			middleware.RestrictToRoles("customer"),
		)).Methods("GET")
}
