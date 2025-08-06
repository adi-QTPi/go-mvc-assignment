package api

import (
	api "github.com/adi-QTPi/go-mvc-assignment/pkg/api/routes"
	"github.com/gorilla/mux"
)

func SetupNewRouter() *mux.Router {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	api.ImplimentApiRoutes(apiRouter) //don't import from api, coz same package

	staticRouter := router.PathPrefix("/static").Subrouter()
	api.ImplimentStaticRoutes(staticRouter)

	return router
}
