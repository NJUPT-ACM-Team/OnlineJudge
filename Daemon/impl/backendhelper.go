package impl

import (
	"OnlineJudge/Daemon/utils"
	"OnlineJudge/mq"
	"OnlineJudge/pbgen/rpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type backendHelperServer struct {
	jmq *mq.MQ
}

func NewBackendHelperServer() *backendHelperServer {
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
	go utils.SubmitToMQ(this.jmq, req)
	return &rpc.SubmitCodeResponse{
		Received: true,
	}, nil
}

func RegisterBackendHelper(server *grpc.Server) {
	rpc.RegisterBackendHelperServer(server, NewBackendHelperServer())
}

type BackendHelper struct {
	conn   *grpc.ClientConn
	bind   string
	client rpc.BackendHelperClient
}

func NewBackendHelper() *BackendHelper {
	return &BackendHelper{
		bind: BIND,
	}
}

func (this *BackendHelper) Connect() error {
	var err error
	this.conn, err = grpc.Dial(this.bind, grpc.WithInsecure())
	return err
}

func (this *BackendHelper) Disconnect() {
	this.conn.Close()
}

func (this *BackendHelper) Submit(run_id int64) (*rpc.SubmitCodeResponse, error) {
	req := &rpc.SubmitCodeRequest{
		RunId: run_id,
	}
	return this.client.Submit(context.Background(), req)

}

func (this *BackendHelper) NewClient() {
	this.client = rpc.NewBackendHelperClient(this.conn)
}
