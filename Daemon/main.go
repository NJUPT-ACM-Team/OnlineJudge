package main

import (
	"OnlineJudge/Daemon/irpc"
	"OnlineJudge/db"
	"OnlineJudge/mq"

	"fmt"
)

func init() {
	db.Init()
	mq.Init()
	irpc.Init()
}

func main() {
	fmt.Println("Hello Daemon.")
	irpc.Run()
}
