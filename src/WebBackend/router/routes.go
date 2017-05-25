package router

import (
	"OnlineJudge/base"
	"OnlineJudge/logger"
	"WebBackend/controller"

	"github.com/gorilla/mux"

	// "log"
	"net/http"
	"time"
)

var log = logger.GetWebBackendLogger()

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

		log.Infof(
			"%s %s %s %s %s",
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
		Name:        "Captcha",
		Method:      "GET",
		Pattern:     "/captcha",
		HandlerFunc: ctler.Captcha,
	},
	Route{
		Name:        "About",
		Method:      "GET",
		Pattern:     "/about",
		HandlerFunc: ctler.About,
	},
	Route{
		Name:        "Register",
		Method:      "POST",
		Pattern:     "/register",
		HandlerFunc: ctler.Register,
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
		Name:        "Logout",
		Method:      "POST",
		Pattern:     "/logout",
		HandlerFunc: ctler.Logout,
	},
	Route{
		Name:        "ListProblems",
		Method:      "GET",
		Pattern:     "/problems",
		HandlerFunc: ctler.ListProblems,
	},
	Route{
		Name:        "ShowProblem",
		Method:      "GET",
		Pattern:     "/problem",
		HandlerFunc: ctler.ShowProblem,
	},
	Route{
		Name:        "Submit",
		Method:      "POST",
		Pattern:     "/submit",
		HandlerFunc: ctler.Submit,
	},
	Route{
		Name:        "ListSubmissions",
		Method:      "GET",
		Pattern:     "/status",
		HandlerFunc: ctler.ListSubmissions,
	},
	Route{
		Name:        "ListContests",
		Method:      "GET",
		Pattern:     "/contests",
		HandlerFunc: ctler.ListContests,
	},
	Route{
		Name:        "ContestShow",
		Method:      "GET",
		Pattern:     "/contest",
		HandlerFunc: ctler.ContestShow,
	},
	Route{
		Name:        "ContestSave",
		Method:      "POST",
		Pattern:     "/contest_save",
		HandlerFunc: ctler.ContestSave,
	},
	Route{
		Name:        "ContestAuth",
		Method:      "POST",
		Pattern:     "/contest_auth",
		HandlerFunc: ctler.ContestAuth,
	},
	Route{
		Name:        "ContestListProblems",
		Method:      "GET",
		Pattern:     "/contest_problems",
		HandlerFunc: ctler.ContestListProblems,
	},
	Route{
		Name:        "ContestSubmit",
		Method:      "POST",
		Pattern:     "/contest_submit",
		HandlerFunc: ctler.ContestSubmit,
	},
	Route{
		Name:        "ContestShowProblem",
		Method:      "GET",
		Pattern:     "/contest_problem",
		HandlerFunc: ctler.ContestShowProblem,
	},
}
