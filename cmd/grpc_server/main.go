package main

import (
	"context"
	"errors"
	"github.com/neofelisho/go-micro-service/config"
	"github.com/neofelisho/go-micro-service/pkg/database"
	"github.com/neofelisho/go-micro-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
	"os"
)

type greeterServer struct {
	proto.UnimplementedGreeterServer
}

func (g greeterServer) SayHello(_ context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	if request == nil {
		return nil, errors.New("nil request")
	}
	log.Printf("LOG: Say hello to: %v", request)

	msg, err := database.SayHello(request.Name)
	if err != nil {
		return nil, err
	}
	return &proto.HelloReply{Message: msg}, nil
}

func main() {
	listener, err := net.Listen("tcp", config.MustLoad().GRPC.BindingAddress())
	if err != nil {
		log.Fatalf("failed to create tcp connection: %v", err)
	}

	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stderr))
	server := grpc.NewServer()
	proto.RegisterGreeterServer(server, &greeterServer{})

	log.Printf("start gRPC server at: %s", config.MustLoad().GRPC.BindingAddress())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
