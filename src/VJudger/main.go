package main

import (
	"OnlineJudge/config"
	"OnlineJudge/judger"
	"OnlineJudge/mq"
	"VJudger/vjudger"
)

func init() {
	cfg := config.Load("")
	mq.Init(cfg.GetMQConfig())
	judger.Init()
}

func main() {
	judger.RunVJ(vjudger.EntryPoint)
}
