package routers

import (
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	router := mux.NewRouter()
	RegisterAll(router)
	return router
}

func RegisterAll(router *mux.Router) {
	inline := router.PathPrefix("/api/inline").Subrouter()
	RegisterInline(inline)

	open := router.PathPrefix("/api/open").Subrouter()
	RegisterOpen(open)
}

func RegisterInline(router *mux.Router) {
	// Register inline login API
	RegisterCommonAPIs(router)
}

func RegisterOpen(router *mux.Router) {
	// Register open login API
	RegisterCommonAPIs(router)
}

func RegisterCommonAPIs(router *mux.Router) {

}
