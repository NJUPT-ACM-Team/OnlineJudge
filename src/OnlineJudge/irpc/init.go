package irpc

import (
	"google.golang.org/grpc"

	"log"
	"net"
)

var BIND string

func Init() {
	BIND = "192.168.56.1:9999"
}

func Run() {
	lis, err := net.Listen("tcp", BIND)
	if err != nil {
		log.Panic(err)
	}
	server := grpc.NewServer()

	RegisterHelper(server)

	server.Serve(lis)
}
