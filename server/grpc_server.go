package server

import (
	"context"
	"fmt"
	"github.com/guobingithub/graceful-shut/constants"
	"github.com/guobingithub/graceful-shut/graceful"
	"github.com/guobingithub/graceful-shut/logger"
	pb "github.com/guobingithub/graceful-shut/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:5555"
)

type HelloService struct {
}

func (r *HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "Hello " + in.Name + "."

	return resp, nil
}

func registerGRpcServer(s *grpc.Server) {
	pb.RegisterHelloServer(s, &HelloService{})
}

func StartGrpcServer(sigSvr *graceful.SignalServer) {
	conn, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("StartGrpcServer, failed to listen: %v", err)
	}

	grpcSvr := grpc.NewServer()
	registerGRpcServer(grpcSvr)

	println("StartGrpcServer, Grpc server Listen on ------" + Address)

	sigSvr.RegisterGracefulShutObj(&graceful.GracefulShutObj{
		Name: constants.GRPC_SERVER,
		Func: func() {
			grpcSvr.GracefulStop()
		},
	})

	if err = grpcSvr.Serve(conn); err != nil {
		logger.Error(fmt.Sprintf("StartGrpcServer, fail to listen!"))
		conn.Close()
	}
}
