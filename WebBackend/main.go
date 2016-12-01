package main

import (
	"OnlineJudge/WebBackend/routers"

	"github.com/gorilla/context"

	"net/http"
)

func main() {
	router := routers.Init()
	// http.Handle("/", router)
	http.ListenAndServe(":8000", context.ClearHandler(router))
}
