package controller

import (
	"context"
	"grpc-boot-starter/apis/protogen"
	"log"
)

func NewHelloControllerImpl() *HelloControllerImpl {
	return &HelloControllerImpl{}
}

// implementing HelloService
type HelloControllerImpl struct {
	protogen.UnimplementedHelloControllerServer
}

func (s *HelloControllerImpl) SayHello(ctx context.Context, in *protogen.HelloRequest) (*protogen.HelloResponse, error) {
	log.Printf("Income request. %v\n", in)
	return &protogen.HelloResponse{Message: "Hello " + in.Name}, nil
}
