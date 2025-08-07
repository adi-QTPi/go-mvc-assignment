package api

import (
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/controllers"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/middleware"
	"github.com/gorilla/mux"
)

func ImplimentAccountRoutes(subRouter *mux.Router) {
	accountController := controllers.NewAccountController()

	subRouter.Handle("/signup",
		middleware.Chain(
			http.HandlerFunc(accountController.CreateNewUser),
			middleware.HashPasword,
			middleware.VerifyDuplicatePassword,
		)).Methods("POST")

	subRouter.Handle("/login",
		middleware.Chain(
			http.HandlerFunc(accountController.LogUserIn),
			middleware.CheckPassword,
			middleware.CheckIfUserExists,
		)).Methods("POST")
}
