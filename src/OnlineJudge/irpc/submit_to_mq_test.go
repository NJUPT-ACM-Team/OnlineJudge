package irpc

import (
	"OnlineJudge/db"
	"OnlineJudge/mq"
	"OnlineJudge/pbgen/rpc"

	"testing"
)

func TestSubmitCodeToMQ(t *testing.T) {
	mq.Init()
	db.Init()
	jmq := mq.New()
	if err := jmq.Connect(); err != nil {
		t.Fatal(err)
	}
	if err := jmq.DeclareVJ(); err != nil {
		t.Fatal(err)
	}
	req := &rpc.StartJudgingRequest{
		RunId: 1,
	}
	SubmitToMQ(jmq, req)
}
