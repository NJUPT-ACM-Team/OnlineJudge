package main

import (
	"OnlineJudge/mq"
	"log"
	"time"
)

func handle(body []byte) {
	log.Println("Recevied: ", string(body))
	time.Sleep(2 * time.Second)
}

func main() {
	forever := make(chan bool)
	mq.Init()
	m := mq.New()
	if err := m.Connect(); err != nil {
		log.Fatal(err)
		return
	}
	defer m.Disconnect()
	if err := m.DeclareVJ(); err != nil {
		log.Fatal(err)
		return
	}
	if err := m.VJReceiver(handle); err != nil {
		log.Fatal(err)
		return
	}
	<-forever
}
