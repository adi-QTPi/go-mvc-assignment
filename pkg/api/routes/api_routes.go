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
	// subRouter.Use(testMiddleware)

	userController := controllers.NewUserApiController()

	subRouter.Handle("/users",
		middleware.Chain(
			http.HandlerFunc(userController.GetAllUsers),
			middleware.RequiredEntries("id", "username", "password"),
			middleware.AnotherMiddleware,
		)).Methods("GET")

	subRouter.HandleFunc("/user/{id}", userController.GetUserById).Methods("GET")

}
