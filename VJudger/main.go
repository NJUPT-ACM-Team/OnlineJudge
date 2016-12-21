package main

import (
	"OnlineJudge/VJudger/vjudger"
	"OnlineJudge/config"
	"OnlineJudge/judger"
	"OnlineJudge/mq"
)

func init() {
	cfg := config.Load("")
	mq.Init(cfg.GetMQConfig())
	judger.Init()
}

func main() {
	judger.RunVJ(vjudger.EntryPoint)
}
