package impl

import (
	"OnlineJudge/Daemon/pb"
	//	"OnlineJudge/models"
	//	"OnlineJudge/models/db"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type backendHelperServer struct{}

func (this *backendHelperServer) Submit(ctx context.Context, req *pb.SubmitRequest) (*pb.SubmitResponse, error) {
	// Submit the code to MQ
	go SumitCode(req)
	return &pb.SubmitResponse{
		Received: true,
	}, nil
}

func RegisterBackendHelper(server *grpc.Server) {
	pb.RegisterBackendHelperServer(server, &backendHelperServer{})
}
