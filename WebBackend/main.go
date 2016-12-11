package main

import (
	"OnlineJudge/Daemon/irpc"
	"OnlineJudge/WebBackend/router"

	"github.com/gorilla/context"

	"net/http"
)

func Init() {
	irpc.Init()
}

func main() {
	Init()
	router := router.Init()
	// http.Handle("/", router)
	http.ListenAndServe(":8000", context.ClearHandler(router))
}
