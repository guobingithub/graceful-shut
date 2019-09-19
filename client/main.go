package main

import (
	"context"
	"fmt"
	"github.com/guobingithub/graceful-shut/logger"
	pb "github.com/guobingithub/graceful-shut/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"sync"
	"time"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:5555"
)

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	var wg = new(sync.WaitGroup)
	start := time.Now()
	for j := 0; j < 10000; j++ {
		wg.Add(1)

		time.Sleep(1 * time.Second)
		go func(conn *grpc.ClientConn, wg *sync.WaitGroup) {
			defer wg.Done()

			c := pb.NewHelloClient(conn)
			reqBody := new(pb.HelloRequest)
			reqBody.Name = "gRPC"
			r, err := c.SayHello(context.Background(), reqBody)
			if err != nil {
				logger.Error(fmt.Sprintf("call gRPC SayHello err:(%v)", err))
			} else {
				println("Grpc client get message after call server ------" + r.Message)
			}
		}(conn, wg)
	}

	wg.Wait()
	logger.Error(fmt.Sprintf("grpc client run ok, time:(%v)", time.Since(start)))
}
