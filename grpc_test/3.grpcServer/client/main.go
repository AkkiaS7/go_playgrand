package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcServer/pb/person"
	"sync"
	"time"
)

func main() {
	l, _ := grpc.Dial("localhost:8888", grpc.WithInsecure())
	client := person.NewSearchServiceClient(l)

	// ping-pong
	res, err := client.Search(context.Background(), &person.PersonReq{Name: "张三"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	// 流式入参
	cIn, _ := client.SearchIn(context.Background())
	for i := 0; i < 3; i++ {
		cIn.Send(&person.PersonReq{Name: "我是进来的信息"})
		time.Sleep(1 * time.Second)
	}
	res, _ = cIn.CloseAndRecv()
	fmt.Println(res)

	// 流式出参
	cOut, _ := client.SearchOut(context.Background(), &person.PersonReq{Name: "李四"})
	for {
		res, err := cOut.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(res)
	}

	// 流式IO
	cIO, _ := client.SearchIO(context.Background())
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(2)
		go func() {
			cIO.Send(&person.PersonReq{Name: "王五"})
			defer wg.Done()
		}()
		go func() {
			res, err := cIO.Recv()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(res)
			defer wg.Done()
		}()
		time.Sleep(1 * time.Second)
		wg.Wait()
	}

}
