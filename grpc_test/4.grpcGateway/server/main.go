package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"grpcGateway/pb/person"
	"net"
	"net/http"
)

type PersonServer struct {
	person.UnimplementedSearchServiceServer
}

func (*PersonServer) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{
		Name: "我收到了" + name + "的信息, 来自grpcGateWay",
	}
	return res, nil
}
func main() {
	go registerGateway()
	go registerGRPC()
	select {}
}

func registerGateway() {
	conn, _ := grpc.DialContext(
		context.Background(),
		"localhost:8888",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)

	mux := runtime.NewServeMux()

	gwServer := &http.Server{
		Handler: mux,
		Addr:    ":8090",
	}

	_ = person.RegisterSearchServiceHandler(context.Background(), mux, conn)
	gwServer.ListenAndServe()
}

func registerGRPC() {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &PersonServer{})
	s.Serve(l)

}
