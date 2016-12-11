package main

import (
	"OnlineJudge/Daemon/impl"
	"OnlineJudge/WebBackend/router"

	"github.com/gorilla/context"

	"net/http"
)

func Init() {
	impl.Init()
}

func main() {
	Init()
	router := router.Init()
	// http.Handle("/", router)
	http.ListenAndServe(":8000", context.ClearHandler(router))
}
