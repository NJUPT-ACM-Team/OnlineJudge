package main

import (
	"LJudger/ljudger"
	"OnlineJudge/config"
	"OnlineJudge/irpc"
	"OnlineJudge/judger"
	"OnlineJudge/mq"
)

func init() {
	cfg := config.Load("")
	mq.Init(cfg.GetMQConfig())
	irpc.Init()
	judger.Init()
}

func main() {
	judger.RunLJ(ljudger.EntryPoint)
}
