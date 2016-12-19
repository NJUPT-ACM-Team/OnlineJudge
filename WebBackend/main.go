package main

import (
	"OnlineJudge/Daemon/irpc"
	"OnlineJudge/WebBackend/router"
	"OnlineJudge/db"

	"github.com/gorilla/context"

	"net/http"
)

func init() {
	db.Init()
	irpc.Init()
}

func main() {
	router := router.Init()
	// http.Handle("/", router)
	panic(http.ListenAndServe(":8000", context.ClearHandler(router)))
}
