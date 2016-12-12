package main

import (
	"OnlineJudge/VJudger/vjudger"
	"OnlineJudge/judger"
	"OnlineJudge/mq"
)

func init() {
	mq.Init()
	judger.Init()
}

func main() {
	judger.RunVJ(vjudger.EntryPoint)
}
