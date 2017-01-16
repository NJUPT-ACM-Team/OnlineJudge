package main

import (
	"OnlineJudge/WebBackend/router"
	"OnlineJudge/config"
	"OnlineJudge/db"
	"OnlineJudge/irpc"

	"github.com/gorilla/context"

	"net/http"
)

func init() {
	cfg := config.Load("")
	db.Init(cfg.GetDBConfig())
	irpc.Init()
}

func main() {
	router := router.Init()
	// http.Handle("/", router)
	panic(http.ListenAndServe(":8000", context.ClearHandler(router)))
}
