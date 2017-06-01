package main

import (
	"LJudger/ljudger"
	"OnlineJudge/config"
	"OnlineJudge/irpc"
	"OnlineJudge/judger"
	"OnlineJudge/mq"

	"flag"
)

type Options struct {
	CfgDir string
}

func MustParseArgs() *Options {
	cfgdir := flag.String(
		"c",
		"config.json",
		"path of config file")
	flag.Parse()
	opts := &Options{
		CfgDir: *cfgdir,
	}
	return opts
}

func init() {
	opts := MustParseArgs()
	cfg := config.Load(opts.CfgDir)
	mq.Init(cfg.GetMQConfig())
	irpc.Init()
	// judger.Init()
}

func main() {
	judger.RunLJ(ljudger.EntryPoint)
}
