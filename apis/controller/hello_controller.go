package controller

import (
	"context"
	pb "grpc-boot-starter/apis/protogen/hello/v1"
	"log"
)

func NewHelloControllerImpl() *HelloControllerImpl {
	return &HelloControllerImpl{}
}

// implementing HelloService
type HelloControllerImpl struct {
	pb.UnimplementedHelloControllerServer
}

func (s *HelloControllerImpl) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	log.Printf("Income request. %v\n", in)
	return &pb.SayHelloResponse{Message: "Hello " + in.Name}, nil
}
