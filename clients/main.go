package main

import (
	"context"
	"fmt"
	"grpc-boot-starter/protogen"
	"log"
	"path/filepath"
	"runtime"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
)

// basepath is the root directory of this package.
var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

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
	}
	// create a new gRPC client
	address := fmt.Sprintf("%s:///unused", r.Scheme())
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
		AccessToken: "some-secret-token",
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
