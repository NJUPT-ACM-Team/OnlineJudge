package main

import (
	"OnlineJudge/Daemon/impl"
	"OnlineJudge/db"
	"OnlineJudge/mq"

	"fmt"
)

func Init() {
	db.Init()
	mq.Init()
	impl.Init()
}

func main() {
	fmt.Println("Hello Daemon.")
	Init()
	impl.Run()
}
