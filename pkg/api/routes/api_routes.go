package api

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/controllers"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/middleware"
	"github.com/gorilla/mux"
)

func testMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("this is the first middleware")

		next.ServeHTTP(w, r)
	})
}

func ImplimentApiRoutes(subRouter *mux.Router) {
	userController := controllers.NewUserApiController()

	subRouter.Handle("/signup",
		middleware.Chain(
			http.HandlerFunc(userController.AddUser),
		)).Methods("POST")

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

}
