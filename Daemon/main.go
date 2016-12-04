package main

import (
	"OnlineJudge/Daemon/impl"

	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Hello Daemon.")
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := impl.Init()
	grpcServer.Serve(lis)
}
