package router

import (
	"OnlineJudge/WebBackend/controller"

	"github.com/gorilla/mux"

	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func RegisterHandlers(router *mux.Router, routes Routes) {
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}

var ctler = controller.NewController(true)

var api_routes = Routes{
	Route{
		Name:        "HelloWorld",
		Method:      "GET",
		Pattern:     "/helloworld",
		HandlerFunc: ctler.HelloWorld,
	},
	Route{
		Name:        "LoginInit",
		Method:      "POST",
		Pattern:     "/login/init",
		HandlerFunc: ctler.LoginInit,
	},
	Route{
		Name:        "LoginAuth",
		Method:      "POST",
		Pattern:     "/login/auth",
		HandlerFunc: ctler.LoginAuth,
	},
}
