package vjudger

import (
	"OnlineJudge/irpc"
	"OnlineJudge/judger"

	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// manual judge
type Result struct {
	Status     string
	StatusCode string
}

func ManualJudge(oj string, pid string, src string, lang string) *Result {
	time.Sleep(10 * time.Second)
	fmt.Printf("Problem Sid: %s-%s\n", oj, pid)
	fmt.Printf("Language:%s\nCode:\n%s\n", lang, src)
	fmt.Printf("1.ac\n2.wa\n3.panic\nchoice:")
	var in int
	// fmt.Scanf("%d", &in)
	in = rand.Intn(2) + 1
	fmt.Println(in)
	switch in {
	case 1:
		return &Result{Status: "Accepted", StatusCode: "ac"}
	case 2:
		return &Result{Status: "Wrong Answer", StatusCode: "wa"}
	case 3:
		panic(errors.New("schedule panic"))
	}
	return &Result{Status: "System Error", StatusCode: "se"}
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

	// Set judging
	res, err := helper.UpdateSubmissionStatus(&irpc.SubmissionStatus{RunId: jdi.GetRunId(), Status: "Judging", StatusCode: "wt"})
	if err != nil {
		log.Println(err)
	}
	log.Println(res)

	// Use manual judge for demo
	j_res := ManualJudge(jdi.GetOJName(), jdi.GetOJPid(), jdi.GetCode(), jdi.GetLanguage().GetLang())
	subs := &irpc.SubmissionStatus{RunId: jdi.GetRunId(), Status: j_res.Status, StatusCode: j_res.StatusCode}
	log.Println(subs)

	res, err = helper.UpdateSubmissionStatus(subs)
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
