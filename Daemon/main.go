package main

import (
	"OnlineJudge/Daemon/irpc"
	"OnlineJudge/db"
	"OnlineJudge/mq"

	"fmt"
)

func Init() {
	db.Init()
	mq.Init()
	irpc.Init()
}

func main() {
	fmt.Println("Hello Daemon.")
	Init()
	irpc.Run()
}
