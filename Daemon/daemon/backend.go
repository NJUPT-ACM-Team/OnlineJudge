package daemon

import (
	"golang.org/x/net/context"
)

type backendHelperServer struct {
}

func (this *backendHelperServer) Submit(ctx context.Context, req *SubmitRequest) (*SubmitResponse, error) {
	// Submit the code to MQ
	return &SubmitResponse{
		Received: true,
	}, nil
}
