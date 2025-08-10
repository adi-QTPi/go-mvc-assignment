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
			middleware.RequiredEntries("user_name", "name", "pwd", "re_pwd", "role"),
		)).Methods("POST")

	subRouter.Handle("/login",
		middleware.Chain(
			http.HandlerFunc(accountController.LogUserIn),
			middleware.CheckPassword,
			middleware.CheckIfUserExists,
			middleware.RequiredEntries("user_name", "password"),
		)).Methods("POST")

	subRouter.Handle("/logout",
		middleware.Chain(
			http.HandlerFunc(accountController.LogUserOut),
		)).Methods("POST")

	subRouter.Handle("/",
		middleware.Chain(
			http.HandlerFunc(accountController.ShowProfile),
			middleware.IdentifyUser,
		)).Methods("GET")

}
