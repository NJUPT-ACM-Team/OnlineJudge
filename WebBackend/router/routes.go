package router

import (
	"OnlineJudge/WebBackend/controller"
	"OnlineJudge/base"

	"github.com/gorilla/mux"

	"log"
	"net/http"
	"time"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s\t%s",
			base.GetIPAddress(r),
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func RegisterHandlers(router *mux.Router, routes Routes) {
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
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
	Route{
		Name:        "ShowProblem",
		Method:      "GET",
		Pattern:     "/problem",
		HandlerFunc: ctler.ShowProblem,
	},
}
