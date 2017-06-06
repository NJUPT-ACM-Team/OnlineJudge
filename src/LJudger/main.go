package main

import (
	"LJudger/ljudger"
	"OnlineJudge/config"
	"OnlineJudge/irpc"
	"OnlineJudge/judger"
	"OnlineJudge/mq"

	"flag"
	"log"
)

type LJudgerConfig struct {
	CorePath  string
	JudgeRoot string
}

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
	log.Println(cfg)
	mq.Init(cfg.GetMQConfig())
	irpc.Init(cfg.GetIRPCConfig())
	ljudger.Init(cfg.LJudger.CorePath, cfg.LJudger.JudgeRoot, cfg.LJudger.UseSudo)
}

func main() {
	judger.RunLJ(ljudger.EntryPoint)
}
