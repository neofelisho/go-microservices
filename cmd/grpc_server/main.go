package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/neofelisho/go-micro-service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type greeterServer struct {
	proto.UnimplementedGreeterServer
}

func (g greeterServer) SayHello(_ context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	if request == nil {
		return nil, errors.New("nil request")
	}
	return &proto.HelloReply{Message: fmt.Sprintf("hello, %v", request.Name)}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to create tcp connection: %v", err)
	}
	server := grpc.NewServer()
	proto.RegisterGreeterServer(server, &greeterServer{})
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
