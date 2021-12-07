package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"mxshop/usersrv/global"
	"mxshop/usersrv/handler"
	"mxshop/usersrv/initialize"
	"mxshop/usersrv/proto"
	"net"
)

func main() {
	port, err := initialize.InitPort()
	if err != nil {
		panic(err.Error())
	}
	err = initialize.InitConfig()
	if err != nil {
		panic("config error" + err.Error())
	}
	global.InitMysql()
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	addr := fmt.Sprintf("%s:%d", "0.0.0.0", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err.Error())
	}
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	err = initialize.ConsulRegister(port, "127.0.0.1")
	if err != nil {
		panic("consul注册失败" + err.Error())
	}
	err = server.Serve(lis)
	if err != nil {
		panic(err.Error())
	}
}
