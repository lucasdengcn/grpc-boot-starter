package services

import (
	"context"
	pb "grpc-boot-starter/protogen"
	"log"
)

func NewHelloServiceServerImpl() *HelloServiceServerImpl {
	return &HelloServiceServerImpl{}
}

// implementing HelloService
type HelloServiceServerImpl struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloServiceServerImpl) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Income request. %v\n", in)
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}
