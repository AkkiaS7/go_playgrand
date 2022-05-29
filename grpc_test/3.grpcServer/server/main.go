package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcServer/pb/person"
	"net"
	"strconv"
	"time"
)

type PersonServer struct {
	person.UnimplementedSearchServiceServer
}

func (*PersonServer) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{
		Name: "我收到了" + name + "的信息",
	}
	return res, nil
}
func (*PersonServer) SearchIn(server person.SearchService_SearchInServer) error {
	for {
		req, err := server.Recv()
		fmt.Println(req)
		if err != nil {
			server.SendAndClose(&person.PersonRes{
				Name: "完成了",
			})
			break
		}
	}
	return nil
}
func (*PersonServer) SearchOut(req *person.PersonReq, server person.SearchService_SearchOutServer) error {
	name := req.Name
	for i := 0; i < 3; i++ {
		server.Send(&person.PersonRes{
			Name: name + "的第" + strconv.Itoa(i) + "条回复信息",
		})
		time.Sleep(1 * time.Second)
	}

	return nil
}
func (*PersonServer) SearchIO(server person.SearchService_SearchIOServer) error {
	str := make(chan string)
	go func() {
		for {
			req, err := server.Recv()

			if err != nil {
				fmt.Println(err)
				str <- "结束"
				break
			}
			str <- req.Name
			fmt.Println(req.Name)
		}
	}()
	for {
		s := <-str
		if s == "结束" {
			break
		}
		server.Send(&person.PersonRes{
			Name: s,
		})
	}
	return nil
}

func main() {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &PersonServer{})
	s.Serve(l)
}
