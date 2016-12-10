package impl

import (
	"OnlineJudge/mq"
	"OnlineJudge/pbgen/rpc"
	//	"OnlineJudge/models"
	//	"OnlineJudge/models/db"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type backendHelperServer struct {
	jmq *mq.MQ
}

func NewBackendHelperServer() *backendHelperServer {
	mq.Init()
	jmq := mq.New()
	if err := jmq.Connect(); err != nil {
		panic(err)
	}
	if err := jmq.DeclareLJ(); err != nil {
		panic(err)
	}
	if err := jmq.DeclareVJ(); err != nil {
		panic(err)
	}
	return &backendHelperServer{
		jmq: jmq,
	}
}

func (this *backendHelperServer) Submit(ctx context.Context, req *rpc.SubmitCodeRequest) (*rpc.SubmitCodeResponse, error) {
	// Submit the code to MQ
	// go SumitCode(req)
	return &rpc.SubmitCodeResponse{
		Received: true,
	}, nil
}

func RegisterBackendHelper(server *grpc.Server) {
	rpc.RegisterBackendHelperServer(server, NewBackendHelperServer())
}