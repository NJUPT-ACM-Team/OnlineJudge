package main

import (
	"OnlineJudge/config"
	"OnlineJudge/db"
	"OnlineJudge/irpc"
	"OnlineJudge/mq"

	"flag"
	"fmt"
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
		"b",
		"127.0.0.1:9999",
		"bind address and port")
	flag.Parse()
	opts := &Options{
		CfgDir: *cfgdir,
		Addr:   *addr,
	}
	return opts
}

func init() {
	opts := MustParseArgs()
	cfg := config.Load(opts.CfgDir)
	db.Init(cfg.GetDBConfig())
	mq.Init(cfg.GetMQConfig())
	irpc.Init()
}

func main() {
	fmt.Println("Hello Daemon.")
	irpc.Run()
}
