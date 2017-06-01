package irpc

import (
	"google.golang.org/grpc"

	"log"
	"net"
)

type IRPCConfig struct {
	ServerBindAddress    string
	ClientConnectAddress string
}

var BIND, CONNECT string

func Init(cfg *IRPCConfig) {
	BIND = "192.168.56.1:9999"
	BIND = cfg.ServerBindAddress
	CONNECT = cfg.ClientConnectAddress
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
