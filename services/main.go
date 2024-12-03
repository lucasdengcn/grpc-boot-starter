package main

import (
	"context"
	pb "grpc-example/protogen"
	"log"
	"net"

	"google.golang.org/grpc"
)

// implementing Service

type helloServiceServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *helloServiceServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Income request. %v\n", in)
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}

// starting the server

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &helloServiceServer{})
	grpcServer.Serve(lis)
}
