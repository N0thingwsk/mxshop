package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mxshop/usersrv/proto"
	"testing"
)

func TestName(t *testing.T) {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err.Error())
		}
	}(conn)
	c := proto.NewUserClient(conn)
	r, err := c.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    0,
		PSize: 0,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, x := range r.Data {
		fmt.Println(x)
	}

}
