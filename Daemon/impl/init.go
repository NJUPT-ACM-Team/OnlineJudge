package impl

import (
	"google.golang.org/grpc"

	"log"
	"net"
)

var BIND string

func Init() {
	BIND = "localhost:9999"
}

func Run() {
	lis, err := net.Listen("tcp", BIND)
	if err != nil {
		log.Panic(err)
	}
	server := grpc.NewServer()

	RegisterBackendHelper(server)

	server.Serve(lis)
}
