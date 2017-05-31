package vjudger

import (
	"OnlineJudge/judger"

	"log"
)

// manual judge
type Result struct {
	Status     string
	StatusCode string
}

func EntryPoint(jdi judger.JudgerInterface) {
	log.Println(jdi.GetRunId())
	log.Println(jdi.GetCode())
	RunVJ(jdi)
}

var VJs = []VJudger{&HDUJudger{}, &PKUJudger{}}

func RunVJ(jdi judger.JudgerInterface) {
	for _, vj := range VJs {
		if vj.Match(jdi.GetOJName()) { //init?match
			if err := vj.Run(jdi); err != nil {
				jdi.SetSystemError()
			}
			break
		}
	}
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
