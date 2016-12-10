package impl

import (
	"OnlineJudge/pbgen/rpc"

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

	client := rpc.NewBackendHelperClient(conn)

	req := &rpc.SubmitCodeRequest{
		RunId: 123,
	}
	res, err := client.Submit(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
