package api

import (
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/controllers"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/middleware"
	"github.com/gorilla/mux"
)

func ImplimentApiRoutes(subRouter *mux.Router) {
	userController := controllers.NewUserApiController()

	subRouter.Use(middleware.IdentifyUser)

	subRouter.Handle("/users",
		middleware.Chain(
			http.HandlerFunc(userController.GetUsers),
			middleware.AnotherMiddleware,
		)).Methods("GET")

	subRouter.Handle("/user/{id}",
		middleware.Chain(
			http.HandlerFunc(userController.GetUser),
		)).Methods("GET")

	subRouter.Handle("/user/{id}",
		middleware.Chain(
			http.HandlerFunc(userController.DeleteUser),
		)).Methods("DELETE")

}
