package api

import (
	api "github.com/adi-QTPi/go-mvc-assignment/pkg/api/routes"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	api.ImplimentApiRoutes(apiRouter)

	staticRouter := router.PathPrefix("/static").Subrouter()
	api.ImplimentStaticRoutes(staticRouter)

	return router
}
