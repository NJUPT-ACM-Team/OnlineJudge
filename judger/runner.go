package judger

import (
	"OnlineJudge/mq"
	msgs "OnlineJudge/pbgen/messages"

	"github.com/golang/protobuf/proto"
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
	GetTestCasesBrief() []*msgs.TestCase
	GetLanguage() *msgs.SubmitLanguage

	GetSpjCode() *msgs.SpjCode
	GetTestCase(id int64) *msgs.TestCase
	UpdateStatus(string, string, int)
}

var jmq *mq.MQ
var judger *Judger

func Init() {
	mq.Init()
	jmq = mq.New()
	if err := jmq.Connect(); err != nil {
		panic(err)
	}
	judger = NewJudger()
}

func Wrapper(fn func(JudgerInterface)) func([]byte) {
	return func(b []byte) {
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

func Run(oj string, fn func([]byte)) {
	Init()
	switch oj {
	case "l":
		if err := jmq.DeclareLJ(); err != nil {
			panic(err)
		}
		if err := jmq.LJReceiver(fn); err != nil {
			panic(err)
		}
	case "v":
		if err := jmq.DeclareVJ(); err != nil {
			panic(err)
		}
		if err := jmq.VJReceiver(fn); err != nil {
			panic(err)
		}
	}
	forever := make(chan bool)
	<-forever
}
