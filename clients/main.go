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
			{Addr: "localhost:50052"},
		},
	})
	//
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
	helloServiceClient := protogen.NewHelloServiceClient(conn)
	callHelloService(helloServiceClient)
	//
	bookServiceClient := protogen.NewBookServiceClient(conn)
	//
	for i := 0; i < 10; i++ {
		go func() {
			bookInfo := callBookCreateService(bookServiceClient)
			callBookUpdateService(bookServiceClient, bookInfo)
			callBookQueryService(bookServiceClient)
		}()
	}
	//
	for {
		time.Sleep(10 * time.Second)
	}
}

func callHelloService(client protogen.HelloServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.SayHello(ctx, &protogen.HelloRequest{Name: "Lucas"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}

func callBookCreateService(client protogen.BookServiceClient) *protogen.BookInfo {
	// Create a Book
	bookCreateInput := &protogen.BookCreateInput{
		Title:       "Book A",
		Description: "This is Book A",
		Amount:      100,
		Price:       10.0,
		Category:    protogen.BookCategory_BOOK_CATEGORY_JAVA,
		Author: &protogen.Author{
			Name: "Author J",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//
	bookInfo, err := client.CreateBook(ctx, bookCreateInput)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("book create: %v\n", bookInfo)
	return bookInfo
}

func callBookUpdateService(client protogen.BookServiceClient, book *protogen.BookInfo) {
	// Update a Book
	input := &protogen.BookUpdateInput{
		Id:          book.Id,
		Title:       "Math Book C updated",
		Description: "This is Math Book C",
		Amount:      110,
		Price:       12.0,
		Category:    protogen.BookCategory_BOOK_CATEGORY_MATH,
		Author: &protogen.Author{
			Name: "Author J",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//
	resp, err := client.UpdateBook(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("book update: %v\n", resp)
}

func callBookQueryService(client protogen.BookServiceClient) {
	// Update a Book
	input := &protogen.BookQueryInput{
		Status:    protogen.BookStatus_BOOK_STATUS_ACTIVE.Enum(),
		PageSize:  10,
		PageIndex: 1,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//
	resp, err := client.QueryBooks(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("book query: %v\n", resp)
}
