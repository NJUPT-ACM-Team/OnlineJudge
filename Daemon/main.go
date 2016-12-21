package main

import (
	"OnlineJudge/Daemon/irpc"
	"OnlineJudge/config"
	"OnlineJudge/db"
	"OnlineJudge/mq"

	"fmt"
)

func init() {
	cfg := config.Load("")
	db.Init(cfg.GetDBConfig())
	mq.Init(cfg.GetMQConfig())
	irpc.Init()
}

func main() {
	fmt.Println("Hello Daemon.")
	irpc.Run()
}
