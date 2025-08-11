package api

import (
	api "github.com/adi-QTPi/go-mvc-assignment/pkg/api/routes"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	api.ImplementApiRoutes(apiRouter)

	staticRouter := router.PathPrefix("/static").Subrouter()
	api.ImplementStaticRoutes(staticRouter)

	accountRouter := router.PathPrefix("/account").Subrouter()
	api.ImplementAccountRoutes(accountRouter)

	rootRouter := router.PathPrefix("/").Subrouter()
	api.ImplementRootRoutes(rootRouter)

	return router
}
