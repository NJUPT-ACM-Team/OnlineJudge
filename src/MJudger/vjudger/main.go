package main

import (
	"OnlineJudge/config"
	"OnlineJudge/irpc"
	"OnlineJudge/judger"
	"OnlineJudge/mq"
	"VJudger/vjudger"
)

func init() {
	cfg := config.Load("")
	mq.Init(cfg.GetMQConfig())
	irpc.Init()
	judger.Init()
}

func main() {
	judger.RunVJ(vjudger.EntryPoint)
}
