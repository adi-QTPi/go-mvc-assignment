package api

import (
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/controllers"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/middleware"
	"github.com/gorilla/mux"
)

func ImplementApiRoutes(subRouter *mux.Router) {

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
			middleware.RestrictToRoles("admin"),
		)).Methods("DELETE")

	itemController := controllers.NewItemApiController()
	subRouter.Handle("/item",
		middleware.Chain(
			http.HandlerFunc(itemController.GetItems),
		)).Methods("GET")
	subRouter.Handle("/item",
		middleware.Chain(
			http.HandlerFunc(itemController.AddItem),
			middleware.RequiredEntries("item_name", "price", "cat_id"),
			middleware.RestrictToRoles("admin"),
		)).Methods("POST")
	subRouter.Handle("/item/{item_id}",
		middleware.Chain(
			http.HandlerFunc(itemController.DeleteItem),
			middleware.RestrictToRoles("admin"),
		)).Methods("DELETE")

	catController := controllers.NewCatApiController
	subRouter.Handle("/categories",
		middleware.Chain(
			http.HandlerFunc(catController().GetCategories),
		)).Methods("GET")
	subRouter.Handle("/categories",
		middleware.Chain(
			http.HandlerFunc(catController().AddCategory),
			middleware.RequiredEntries("category_name", "category_description"),
			middleware.RestrictToRoles("admin"),
		)).Methods("POST")

	orderController := controllers.NewOrderApiController()
	subRouter.Handle("/order",
		middleware.Chain(
			http.HandlerFunc(orderController.PlaceOrder),
			middleware.AssignEmptyTable,
			middleware.DecodeCartJsonInput,
			middleware.RestrictToRoles("customer"),
		)).Methods("POST")
	subRouter.Handle("/pay",
		middleware.Chain(
			http.HandlerFunc(orderController.OrderPayment),
			middleware.RestrictToRoles("customer"),
		)).Methods("POST")

	cookController := controllers.NewCookApiController()
	subRouter.Handle("/cook",
		middleware.Chain(
			http.HandlerFunc(cookController.ChangeKitchenOrderStatus),
			middleware.RequiredEntries("order_id", "item_id", "is_complete"),
			middleware.RestrictToRoles("cook"),
		)).Methods("POST")
}
