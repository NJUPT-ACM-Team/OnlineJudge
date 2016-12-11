package main

import (
	"OnlineJudge/VJudger/vjudger"
	"OnlineJudge/judger"

	"log"
)

func Init() {
	judger.Init()
}

func main() {
	Init()
	judger.RunVJ(vjudger.EntryPoint)
}
