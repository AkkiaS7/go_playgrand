package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc "hello_world/pb"
	"net"
)

// 取出server
type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

// SayHi 挂载方法
func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "我是从服务端返回的grpc内容"}, nil
}

func main() {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})
	s.Serve(l)
}
