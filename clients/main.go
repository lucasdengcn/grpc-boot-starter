package main

import (
	"context"
	"fmt"
	"grpc-boot-starter/apis/protogen"
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
	helloServiceClient := protogen.NewHelloControllerClient(conn)
	callHelloService(helloServiceClient)
	//
	bookServiceClient := protogen.NewBookControllerClient(conn)
	//
	for i := 0; i < 2; i++ {
		go func() {
			bookInfo := callBookCreateService(bookServiceClient)
			// callBookUpdateService(bookServiceClient, bookInfo)
			// callBookQueryService(bookServiceClient)
			callBookGetService(bookServiceClient, bookInfo.Id)
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
		AccessToken: "eyJhbGciOiJQUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlcyI6WyJVc2VyIl0sImdyb3VwcyI6WyJVc2VyIl0sImlzcyI6Imdpbi1ib290LXN0YXJ0ZXIiLCJzdWIiOiIxIiwiYXVkIjpbInRlc3QiXSwiZXhwIjoxNzMzNDgxNTQ2LCJuYmYiOjE3MzM0NDU1NDYsImlhdCI6MTczMzQ0NTU0NiwianRpIjoiOTAwMTUwOTgzY2QyNGZiMGQ2OTYzZjdkMjhlMTdmNzIifQ.OGx9ngKIOGVyTjTSin0AygaJjWwcArrVcL0uhaxlioQIoNcVfWTNq8Q6Z0J34GHPaaw92zjA90B6ru6rbdFbfAwl7IQsMsUqG2wO5cYDk4GU43QADX-O4zCF1m7qy5l8Sbar80qJDG3PsmGwtaS9q-2E4yMPga0VPOVyXsRLA792EH65yPAkA116cnLToRIpzcATm_MBTgw1lL1GKROMX2svBtaJ626KMO4XMoVRRzo2FbtMAVLv4kMlxWg4VTM8gdGjysWpT5Uj7R9iykk3zoevr834RmoBK8eHfJDsjqayjHdKBq-7sjX8k2l3EHsTro4AkEwCR5Lmf0vsMhaiOYZ4iG2M3UN_uLnvDzT0nDhXqLL22H_FZPijN-l0yMOq3Q2DHB6mEinU5zBdsPf25YrB4bbLcfW3R46OFSw_E5Cc1aGcyGWWmKpyRfEeg8VvZ5N58SXtpaqKhEYkUDz5rW4C60dH29NqFTDVRWuICiOVITBPhuPAem-uwx5jFGexi02Bzfx44CXhf6WJysRcHi5IkHegzSYec5mDUJLGno-HI2PpA1DJ3ojgTD76niqMKGkFFIWn6BJRDiMoOh-hQVcVIwvhvpDqy5aFs_f0ujjj2_c39a-IrIJGNsWqOcHp9GcU_I8kl-3d03BASfJVfDSkItOUii64w1gzeHB_Rnc",
	}
}

func callHelloService(client protogen.HelloControllerClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.SayHello(ctx, &protogen.HelloRequest{Name: "Lucas"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}

func callBookCreateService(client protogen.BookControllerClient) *protogen.BookInfo {
	// Create a Book
	bookCreateInput := &protogen.BookCreateInput{
		Title:       "Book A Long title long title long title long title",
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

func callBookGetService(client protogen.BookControllerClient, id uint32) *protogen.BookInfo {
	// Get a Book
	bookGetInput := &protogen.BookGetInput{
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
	bookInfo, err := client.GetBook(ctx, bookGetInput, grpc.Header(&header), grpc.Trailer(&trailer), grpc.WaitForReady(true))
	if err != nil {
		got := status.Code(err)
		log.Println(got)
		log.Fatal(err)
	}
	log.Printf("book get: %v\n", bookInfo)
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
	return bookInfo
}

func callBookUpdateService(client protogen.BookControllerClient, book *protogen.BookInfo) {
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

func callBookQueryService(client protogen.BookControllerClient) {
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
