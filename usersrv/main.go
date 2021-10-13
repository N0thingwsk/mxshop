package main

import (
	"google.golang.org/grpc"
	"mxshop/usersrv/handler"
	"mxshop/usersrv/proto"
	"net"
)

func main() {
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err.Error())
	}
	err = server.Serve(lis)
	if err != nil {
		panic(err.Error())
	}
}
