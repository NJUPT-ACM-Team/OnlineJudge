package main

import (
	"OnlineJudge/config"
	"OnlineJudge/db"
	"OnlineJudge/irpc"
	"WebBackend/router"

	"github.com/gorilla/context"
	"github.com/rs/cors"

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
	handler := cors.Default().Handler(context.ClearHandler(router))
	panic(http.ListenAndServe(":8000", handler))
}
