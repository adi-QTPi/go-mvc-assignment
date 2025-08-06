package api

import (
	"github.com/adi-QTPi/go-mvc-assignment/pkg/controllers"
	"github.com/gorilla/mux"
)

func ImplimentApiRoutes(subRouter *mux.Router) {

	userController := controllers.NewUserApiController()
	subRouter.HandleFunc("/users", userController.GetAllUsers).Methods("GET")

	subRouter.HandleFunc("/user/{id}", userController.GetUserById).Methods("GET")

}
