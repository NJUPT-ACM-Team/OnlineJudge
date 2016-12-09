package impl

import (
	"OnlineJudge/pbgen/rpc"
	//	"OnlineJudge/models"
	//	"OnlineJudge/models/db"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type backendHelperServer struct{}

func (this *backendHelperServer) Submit(ctx context.Context, req *rpc.SubmitCodeRequest) (*rpc.SubmitCodeResponse, error) {
	// Submit the code to MQ
	// go SumitCode(req)
	return &rpc.SubmitCodeResponse{
		Received: true,
	}, nil
}

func RegisterBackendHelper(server *grpc.Server) {
	rpc.RegisterBackendHelperServer(server, &backendHelperServer{})
}
