package main

import (
	"OnlineJudge/judger"

	"log"
)

func handler(j judger.JudgerInterface) {
	log.Println(j.GetCode())
}

func main() {
	judger.RunLJ(handler)
}
