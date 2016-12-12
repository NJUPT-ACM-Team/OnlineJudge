package vjudger

import (
	"OnlineJudge/judger"
	"log"
)

func EntryPoint(jdi judger.JudgerInterface) {
	log.Println(jdi.GetRunId())
}

type VJudger interface {
	Init(judger.JudgerInterface) error
	Login(judger.JudgerInterface) error
	Submit(judger.JudgerInterface) error
	GetStatus(judger.JudgerInterface) error
	Run(judger.JudgerInterface) error
	Match(string) bool
	// Crawler(string) error
}
