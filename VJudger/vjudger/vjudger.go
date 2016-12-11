package vjudger

import (
	"OnlineJudge/judger"
)

func EntryPoint(jdi judger.JudgerInterface) {

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
