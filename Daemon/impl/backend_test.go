package impl

import (
	"OnlineJudge/Daemon/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"testing"
)

func TestBackend(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure())
	if err != nil {
		t.Log(err)
	}
	defer conn.Close()

	client := pb.NewBackendHelperClient(conn)

	req := &pb.SubmitRequest{
		RunId: 123,
	}
	res, err := client.Submit(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
