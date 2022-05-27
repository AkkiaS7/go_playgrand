package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc "grpc_test/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client := hello_grpc.NewHelloGRPCClient(conn)
	res, err := client.SayHi(context.Background(), &hello_grpc.Req{Message: "我是从客户端发送的grpc内容"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.GetMessage())
}
