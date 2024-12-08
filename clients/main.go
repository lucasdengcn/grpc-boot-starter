package main

import (
	"context"
	"fmt"
	pbbook "grpc-boot-starter/apis/protogen/book/v1"
	pbhello "grpc-boot-starter/apis/protogen/hello/v1"
	"log"
	"path/filepath"
	"runtime"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	_ "google.golang.org/grpc/health"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	"google.golang.org/grpc/status"
)

// basepath is the root directory of this package.
var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

// import health package to enable healch check.
var serviceConfig = `{
	"loadBalancingPolicy": "round_robin",
	"healthCheckConfig": {
		"serviceName": ""
	},
	"methodConfig": [{
		"name": [{"service": "protogen.BookService"}],
		"retryPolicy": {
			"MaxAttempts": 4,
			"InitialBackoff": ".01s",
			"MaxBackoff": ".01s",
			"BackoffMultiplier": 1.0,
			"RetryableStatusCodes": [ "UNAVAILABLE" ]
		}
	}]
}`

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

// start a grpc client
func main() {
	fmt.Println(basepath)
	// name resolver
	r := manual.NewBuilderWithScheme("whatever")
	r.InitialState(resolver.State{
		Addresses: []resolver.Address{
			{Addr: "localhost:50051"},
			{Addr: "localhost:50052"},
		},
	})
	// prepare token
	perRPC := oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(fetchToken())}
	// prepare tls certs
	creds := loadClientTLSCert()
	//
	options := []grpc.DialOption{
		grpc.WithPerRPCCredentials(perRPC),
		grpc.WithTransportCredentials(creds),
		grpc.WithResolvers(r),
		grpc.WithDefaultServiceConfig(serviceConfig),
		grpc.WithKeepaliveParams(kacp),
	}
	// create a new gRPC client
	address := fmt.Sprintf("%s:///unused", r.Scheme())
	conn, err := grpc.NewClient(address, options...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// create a new service client
	helloServiceClient := pbhello.NewHelloControllerServiceClient(conn)
	callHelloService(helloServiceClient)
	//
	bookServiceClient := pbbook.NewBookControllerServiceClient(conn)
	//
	for i := 0; i < 2; i++ {
		go func() {
			createBookResp := callBookCreateService(bookServiceClient)
			callBookUpdateService(bookServiceClient, createBookResp.Book)
			callBookQueryService(bookServiceClient)
			callBookGetService(bookServiceClient, createBookResp.Book.Id)
		}()
	}
	//
	for {
		time.Sleep(10 * time.Second)
	}
}

func loadClientTLSCert() credentials.TransportCredentials {
	path := filepath.Dir(basepath)
	creds, err := credentials.NewClientTLSFromFile(filepath.Join(path, "secrets/x509/ca_cert.pem"), "x.test.example.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	return creds
}

// fetchToken simulates a token lookup and omits the details of proper token
// acquisition. For examples of how to acquire an OAuth2 token, see:
// https://godoc.org/golang.org/x/oauth2
func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "eyJhbGciOiJQUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlcyI6WyJVc2VyIl0sImdyb3VwcyI6WyJVc2VyIl0sImlzcyI6Imdpbi1ib290LXN0YXJ0ZXIiLCJzdWIiOiIxIiwiYXVkIjpbInRlc3QiXSwiZXhwIjoxNzMzNjc3NzQ1LCJuYmYiOjE3MzM2NDE3NDUsImlhdCI6MTczMzY0MTc0NSwianRpIjoiOTAwMTUwOTgzY2QyNGZiMGQ2OTYzZjdkMjhlMTdmNzIifQ.eVVLsZIMCD2xTG7BdPlxuDcVLqEf4Pd5e3edS4u7fch_9i8zCNgRDPBnMb6PWprNawuqlavNUwPo4U_z34zwaspfEd-BGz5NnBhpJDkWJitlf27nqN7IwRlCJdfWsHldWQNvsqp_uUCcvJxUwbs8iDoyWYa6dwbCfPUYuOhFN-hk0aAqYcvyQdutimi0TthFitcJhvTQoOjtptg-U4SXp6rHANnwQeh5c6fNTyIOhcrUUhg9PQ0O4iJlpV98A658vrJ05uOjm0HnJgAwWpsC3BVq_rzH1CGRdbwn2n5S-ajrDLT7jXbZm6iYTdVH1YkoIPhqt9wy8wBakIYBeTNWGZcIp7wFqmDxc3DsT0mLZj0MzvnLum_ZFFNwp4Wzwzdn_5kF1VHjtd6eeIjctJWhXTrl6QP9btV2VFRwCC4ItP_YW__cC2KPTCDBnETVS6SC2hZOhzjrcHw6bjsaT4Aun6mmo5heMGsLcBI3wzrAZAsbMrtB2uPzCC263U2TNkMwbTiMId0beiO6mLmJn9CGRqDbAKzzv5QMJ-fEu-dWPNdNMKIVvGXqw1BdUWoSVAoF8QE2HVODBPL8gptxUEt5st1spK7vRkpkeOJ9xygOpfh0dLiis4K5r6kvi-3zphIhD6A-Vqu98um4VuO9rm-R21HwwguhBQfwcNlrfgLcVRU",
	}
}

func callHelloService(client pbhello.HelloControllerServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.SayHello(ctx, &pbhello.SayHelloRequest{Name: "Lucas"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}

func callBookCreateService(client pbbook.BookControllerServiceClient) *pbbook.CreateBookResponse {
	// Create a Book
	bookCreateInput := &pbbook.CreateBookRequest{
		Title:       "Book A Long title long title long title long title",
		Description: "This is Book A",
		Amount:      100,
		Price:       10.0,
		Category:    pbbook.BookCategory_BOOK_CATEGORY_JAVA,
		Author: &pbbook.Author{
			Name: "Author J",
		},
		Email:    "abc@example.com",
		CoverUrl: "",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//
	resp, err := client.CreateBook(ctx, bookCreateInput)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("book create: %v\n", resp)
	return resp
}

func callBookGetService(client pbbook.BookControllerServiceClient, id uint32) *pbbook.GetBookResponse {
	// Get a Book
	bookGetInput := &pbbook.GetBookRequest{
		Id: id,
	}
	// Create metadata and context.
	md := metadata.Pairs("traceparent", "00-0123456789abcdef0123456789abcdef-0123456789abcdef-0")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	// deadline
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	//
	var header, trailer metadata.MD
	resp, err := client.GetBook(ctx, bookGetInput, grpc.Header(&header), grpc.Trailer(&trailer), grpc.WaitForReady(true))
	if err != nil {
		got := status.Code(err)
		log.Println(got)
		log.Fatal(err)
	}
	log.Printf("book get: %v\n", resp)
	// Get header from server
	fmt.Println("Received headers:")
	for k, v := range header {
		fmt.Printf("%s: %v\n", k, v)
	}

	fmt.Println("Received trailers:")
	for k, v := range trailer {
		fmt.Printf("%s: %v\n", k, v)
	}
	//
	return resp
}

func callBookUpdateService(client pbbook.BookControllerServiceClient, book *pbbook.BookInfo) {
	// Update a Book
	input := &pbbook.UpdateBookRequest{
		Id:          book.Id,
		Title:       "Math Book C updated",
		Description: "This is Math Book C",
		Amount:      110,
		Price:       12.0,
		Category:    pbbook.BookCategory_BOOK_CATEGORY_MATH,
		Author: &pbbook.Author{
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

func callBookQueryService(client pbbook.BookControllerServiceClient) {
	// Update a Book
	input := &pbbook.QueryBooksRequest{
		Status:    pbbook.BookStatus_BOOK_STATUS_ACTIVE.Enum(),
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
