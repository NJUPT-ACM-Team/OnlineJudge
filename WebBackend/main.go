package main

import (
	"OnlineJudge/WebBackend/router"

	"github.com/gorilla/context"

	"net/http"
)

func main() {
	router := router.Init()
	// http.Handle("/", router)
	http.ListenAndServe(":8000", context.ClearHandler(router))
}
