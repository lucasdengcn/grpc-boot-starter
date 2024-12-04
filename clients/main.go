package main

import (
	"context"
	"fmt"
	"grpc-boot-starter/protogen"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
)

var serviceConfig = `{
	"loadBalancingPolicy": "round_robin",
	"healthCheckConfig": {
		"serviceName": ""
	}
}`

// start a grpc client
func main() {
	// name resolver
	r := manual.NewBuilderWithScheme("whatever")
	r.InitialState(resolver.State{
		Addresses: []resolver.Address{
			{Addr: "localhost:50051"},
			{Addr: "localhost:50051"},
		},
	})

	address := fmt.Sprintf("%s:///unused", r.Scheme())
	//
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(r),
		grpc.WithDefaultServiceConfig(serviceConfig),
	}
	// create a new gRPC client
	conn, err := grpc.NewClient(address, options...)
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
	for {
		time.Sleep(time.Second)
	}
}
