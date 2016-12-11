package utils

import (
	"OnlineJudge/Daemon/pb"
	"OnlineJudge/mq"
	msgs "OnlineJudge/pbgen/messages"
)

func SubmitToMQ(jmq *mq.MQ, req *pb.SubmitRequest) {
	send := msgs.SubmitCodeRequest{
		RunId: req.RunId,
	}
}
