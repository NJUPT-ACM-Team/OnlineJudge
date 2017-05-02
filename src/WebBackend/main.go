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
	c := cors.New(cors.Options{
		AllowCredentials: true,
	})

	// Insert the middleware
	handler := c.Handler(context.ClearHandler(router))
	panic(http.ListenAndServe(":8000", handler))
}
