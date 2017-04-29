package irpc

import (
	"OnlineJudge/mq"
	"OnlineJudge/pbgen/rpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type helperServer struct {
	jmq *mq.MQ
}

func NewHelperServer() *helperServer {
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
	return &helperServer{
		jmq: jmq,
	}
}

func (this *helperServer) StartJudging(ctx context.Context, req *rpc.StartJudgingRequest) (*rpc.StartJudgingResponse, error) {
	// Submit the code to MQ
	go SubmitToMQ(this.jmq, req)
	return &rpc.StartJudgingResponse{
		Received: true,
	}, nil
}

func (this *helperServer) GetTestingData(ctx context.Context, req *rpc.GetTestingDataRequest) (*rpc.GetTestingDataResponse, error) {
	return &rpc.GetTestingDataResponse{}, nil
}

func (this *helperServer) Register(ctx context.Context, req *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	return &rpc.RegisterResponse{}, nil
}

func RegisterHelper(server *grpc.Server) {
	rpc.RegisterHelperServer(server, NewHelperServer())
}

type Helper struct {
	conn   *grpc.ClientConn
	bind   string
	client rpc.HelperClient
}

func NewHelper() *Helper {
	return &Helper{
		bind: BIND,
	}
}

func (this *Helper) Connect() error {
	var err error
	this.conn, err = grpc.Dial(this.bind, grpc.WithInsecure())
	return err
}

func (this *Helper) Disconnect() {
	this.conn.Close()
}

func (this *Helper) Submit(run_id int64) (*rpc.StartJudgingResponse, error) {
	req := &rpc.StartJudgingRequest{
		RunId: run_id,
	}
	return this.client.StartJudging(context.Background(), req)

}

func (this *Helper) NewClient() {
	this.client = rpc.NewHelperClient(this.conn)
}
