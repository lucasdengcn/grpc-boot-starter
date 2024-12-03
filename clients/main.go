package main

import (
	"context"
	"grpc-example/protogen"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// start a grpc client
func main() {
	// create a new gRPC client
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// create a new service client
	client := protogen.NewHelloServiceClient(conn)

	// call the service method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.SayHello(ctx, &protogen.HelloRequest{Name: "Lucas"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
