package vjudger

import (
	"OnlineJudge/irpc"
	"OnlineJudge/judger"

	"log"
)

// manual judge
type Result struct {
	Status     string
	StatusCode string
}

func ManualJudge(oj string, pid string, src string, lang string) *Result {
}

func EntryPoint(jdi judger.JudgerInterface) {
	log.Println(jdi.GetRunId())
	log.Println(jdi.GetCode())

	helper := irpc.NewHelper()
	if err := helper.Connect(); err != nil {
		// Log the error
		log.Println(err)
		return
	}
	defer helper.Disconnect()

	helper.NewClient()

	subs := &irpc.SubmissionStatus{RunId: jdi.GetRunId(), Status: "Accepted", StatusCode: "ac"}

	res, err := helper.UpdateSubmissionStatus(&irpc.SubmissionStatus{RunId: jdi.GetRunId(), Status: "Accepted", StatusCode: "ac"})
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
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
