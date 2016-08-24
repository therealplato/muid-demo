package main

import (
	"net"

	"github.com/namely/accounts/server"
	pb "github.com/therealplato/muid-demo/protoheaders"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", s.cfg.ListenAddress)
	if err != nil {
		panic(err)
	}

	s := server.New()
	gs := grpc.NewServer()
	pb.RegisterMuidServer(gs, s)
	<-grpc.Serve(lis)
}
