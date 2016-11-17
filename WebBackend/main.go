package main

import (
	"OnlineJudge/WebBackend/routers"

	"net/http"
)

func main() {
	router := routers.Init()
	// http.Handle("/", router)
	http.ListenAndServe(":8000", router)
}
