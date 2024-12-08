package controller

import (
	"context"
	pb "grpc-boot-starter/apis/protov1"
	"log"
)

func NewHelloControllerImpl() *HelloControllerImpl {
	return &HelloControllerImpl{}
}

// implementing HelloService
type HelloControllerImpl struct {
	pb.UnimplementedHelloControllerServiceServer
}

func (s *HelloControllerImpl) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Income request. %v\n", in)
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}
