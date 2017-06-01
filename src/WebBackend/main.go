package main

import (
	"OnlineJudge/config"
	"OnlineJudge/db"
	"OnlineJudge/irpc"
	"WebBackend/router"

	"github.com/gorilla/context"
	"github.com/rs/cors"

	"flag"
	"net/http"
)

type Options struct {
	CfgDir string
	Addr   string
}

func MustParseArgs() *Options {
	cfgdir := flag.String(
		"c",
		"config.json",
		"path of config file")
	addr := flag.String(
		"s",
		":8000",
		"bind address and port")
	flag.Parse()
	opts := &Options{
		CfgDir: *cfgdir,
		Addr:   *addr,
	}
	return opts
}

var opts *Options

func init() {
	opts = MustParseArgs()
	cfg := config.Load(opts.CfgDir)
	db.Init(cfg.GetDBConfig())
	irpc.Init(cfg.GetIRPCConfig())
}

func main() {
	router := router.Init()
	// http.Handle("/", router)
	c := cors.New(cors.Options{
		AllowCredentials: true,
	})

	// Insert the middleware
	handler := c.Handler(context.ClearHandler(router))
	panic(http.ListenAndServe(opts.Addr, handler))
}
