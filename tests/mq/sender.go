package main

import (
	"OnlineJudge/mq"
	"log"
	"os"
)

func send(data string) error {
	mq.Init()
	m := mq.New()
	if err := m.Connect(); err != nil {
		return err
	}
	defer m.Disconnect()
	if err := m.DeclareLJ(); err != nil {
		return err
	}
	if err := m.PublishLJ([]byte(data)); err != nil {
		return err
	}
	return nil
}

func main() {
	data := os.Args[1]
	if err := send(data); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Sent: %s", data)
	}
}
