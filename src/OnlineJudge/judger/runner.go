package judger

import (
	"OnlineJudge/irpc"
	"OnlineJudge/mq"
	msgs "OnlineJudge/pbgen/messages"

	"github.com/golang/protobuf/proto"

	"log"
	"time"
)

type JudgerInterface interface {
	Init(*msgs.SubmitMQ)

	GetRunId() int64
	GetOJName() string
	GetOJPid() string
	GetCode() string
	IsLocal() bool
	IsVirtual() bool
	IsSpj() bool
	GetTimeLimit() int
	GetMemoryLimit() int
	GetTestCasesBrief() []*msgs.TestCase
	GetLanguage() *msgs.SubmitLanguage
	GetSubmitTime() time.Time

	GetSpjCode() *msgs.SpjCode
	GetTestCase(id int64) *msgs.TestCase
	UpdateResource(int, int) error
	UpdateResult(string, string) error
	UpdateCEInfo(string) error
	UpdateStatus(*irpc.SubmissionStatus) error
	UpdateStatusJudging() error
	SetSystemError() error
}

func Wrapper(fn func(JudgerInterface)) func([]byte) {
	return func(b []byte) {
		judger := NewJudger()
		s := &msgs.SubmitMQ{}
		proto.Unmarshal(b, s)
		judger.Init(s)
		fn(judger)
	}
}

func RunVJ(fn func(JudgerInterface)) {
	Run("v", Wrapper(fn))
}

func RunLJ(fn func(JudgerInterface)) {
	Run("l", Wrapper(fn))
}

func RunMJ(fn func(JudgerInterface)) {
	Run("m", Wrapper(fn))
}

func Run(oj string, fn func([]byte)) {
	jmq := mq.New()
	log.Println("ready to connect")
	if err := jmq.Connect(); err != nil {
		log.Println(err)
		panic(err)
	}
	log.Println("connect successfully")
	switch oj {
	case "l":
		log.Println("ready to declare ljudger")
		if err := jmq.DeclareLJ(); err != nil {
			log.Println(err)
			panic(err)
		}
		log.Println("declare successfully")
		if err := jmq.LJReceiver(fn); err != nil {
			log.Println(err)
			panic(err)
		}
		log.Println("receiver successfully")
	case "v":
		if err := jmq.DeclareVJ(); err != nil {
			panic(err)
		}
		if err := jmq.VJReceiver(fn); err != nil {
			panic(err)
		}
	case "m":
		if err := jmq.DeclareMJ(); err != nil {
			panic(err)
		}
		if err := jmq.MJReceiver(fn); err != nil {
			panic(err)
		}
	}
	forever := make(chan bool)
	<-forever
}
