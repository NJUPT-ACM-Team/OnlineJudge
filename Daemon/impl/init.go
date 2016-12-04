package impl

import (
	"google.golang.org/grpc"
)

func Init() *grpc.Server {
	server := grpc.NewServer()

	RegisterBackendHelper(server)

	return server
}
