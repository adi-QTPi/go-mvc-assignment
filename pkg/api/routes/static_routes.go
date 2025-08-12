package api

import (
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/controllers"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/middleware"
	"github.com/gorilla/mux"
)

func ImplementStaticRoutes(subRouter *mux.Router) {

	subRouter.Use(middleware.IdentifyUser)

	staticController := controllers.NewStaticController()
	subRouter.Handle("/menu",
		middleware.Chain(
			http.HandlerFunc(staticController.RenderMenuPage),
		)).Methods("GET")
	subRouter.Handle("/cart",
		middleware.Chain(
			http.HandlerFunc(staticController.RenderCartPage),
		)).Methods("GET")

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

	customerSaticController := controllers.NewStaticOrderController()
	subRouter.Handle("/order",
		middleware.Chain(
			http.HandlerFunc(customerSaticController.RenderCustOrderPage),
			middleware.RestrictToRoles("customer"),
		)).Methods("GET")
	subRouter.Handle("/order/{id}",
		middleware.Chain(
			http.HandlerFunc(customerSaticController.RenderOrderById),
			middleware.RestrictToRoles("customer", "admin"),
		)).Methods("GET")

	renderer := controllers.NewPageRenderer()
	subRouter.Handle("",
		middleware.Chain(
			http.HandlerFunc(renderer.RenderHomePage),
		)).Methods("GET")
}
