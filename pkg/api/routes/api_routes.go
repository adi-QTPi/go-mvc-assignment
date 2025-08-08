package api

import (
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/controllers"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/middleware"
	"github.com/gorilla/mux"
)

func ImplimentApiRoutes(subRouter *mux.Router) {

	subRouter.Use(middleware.IdentifyUser)

	userController := controllers.NewUserApiController()

	subRouter.Handle("/users",
		middleware.Chain(
			http.HandlerFunc(userController.GetUsers),
		)).Methods("GET")

	subRouter.Handle("/user/{id}",
		middleware.Chain(
			http.HandlerFunc(userController.GetUser),
		)).Methods("GET")

	subRouter.Handle("/user/{id}",
		middleware.Chain(
			http.HandlerFunc(userController.DeleteUser),
		)).Methods("DELETE")

	itemController := controllers.NewItemApiController()

	subRouter.Handle("/item",
		middleware.Chain(
			http.HandlerFunc(itemController.GetItems),
		)).Methods("GET")
	subRouter.Handle("/item",
		middleware.Chain(
			http.HandlerFunc(itemController.AddItem),
		)).Methods("POST")
	subRouter.Handle("/item/delete",
		middleware.Chain(
			http.HandlerFunc(itemController.DeleteItem),
		)).Methods("POST")

	catController := controllers.NewCatApiController
	subRouter.Handle("/categories",
		middleware.Chain(
			http.HandlerFunc(catController().GetCategories),
		)).Methods("GET")
	subRouter.Handle("/categories",
		middleware.Chain(
			http.HandlerFunc(catController().AddCategory),
		)).Methods("POST")

	orderController := controllers.NewOrderApiController()
	subRouter.Handle("/order",
		middleware.Chain(
			http.HandlerFunc(orderController.PlaceOrder),
			middleware.AssignEmptyTable,
			middleware.DecodeCartJsonInput,
		)).Methods("POST")
}
