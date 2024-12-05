package main

import (
	"context"
	"fmt"
	"grpc-boot-starter/protogen"
	"log"
	"math/rand"
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
			// bookInfo := callBookCreateService(bookServiceClient)
			// callBookUpdateService(bookServiceClient, bookInfo)
			// callBookQueryService(bookServiceClient)
			callBookGetService(bookServiceClient, uint32(rand.Intn(1000)))
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
		AccessToken: "eyJhbGciOiJQUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlcyI6WyJVc2VyIl0sImdyb3VwcyI6WyJVc2VyIl0sImlzcyI6Imdpbi1ib290LXN0YXJ0ZXIiLCJzdWIiOiIxIiwiYXVkIjpbInRlc3QiXSwiZXhwIjoxNzMzNDA3NjIwLCJuYmYiOjE3MzMzNzE2MjAsImlhdCI6MTczMzM3MTYyMCwianRpIjoiOTAwMTUwOTgzY2QyNGZiMGQ2OTYzZjdkMjhlMTdmNzIifQ.DuyGnzvjyeZTB5ABkNIs3A5gyKz8RO-4sr1hDZKzbIDnkuZ1qTMvXQouh4Iwv5GQmLptOT5AdXhhwt1aczol2H1w2hez8hc12bjfYUw0Far7p7zfk9-rXdvntBasiqzXwuF4eCyrVLORPoYYFIDnegQv0R-D3_BZ56eWGaZZdkm7XRFjdpfFpsZOvsd5h2CZ7djq6n08AVCnP1JdTTZbPIwtHROEl8VfErETOtxagolLZiS2Ju53KXt9aEGrZcmzeYKJt70eDgXzJq7FiHuXzh5fK9EAxSsYipvo6olf2ssm3y1mpXaMMwKTIWsyfejcHyiodNXRWpO4MC-5zLAhmcHlrAZatWbSr2FrrpCk_KEHqi3fJ0lzjcrwOZulEcOeAanTAI_xnjWx6IPbG9zzSb8W6MAb_Gr1BM2NRYG8dl_C-3Xmuk0zxs45mtoo7mrqpBV_COvfhUdM03aho0qxVdtuA9EEtik53e-IYiYwapG5jEKEjJdu-XguV2k3rRMhhXDGN4LJsTgzmXWLwa21FYjC40gF9ASfrRvgZ_bZDgfJknSWJUg1_obPU5JTF1T1YvQFDiK78eU6nkXab8afH3WKoOQG5aAmECch1nVG1MOVpl7IdbaLQRTOp9zVxiO7daC_PG451Js2t8WMGv9nYlo0X0pBNDusgdbezX6XtM4",
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

func callBookGetService(client protogen.BookServiceClient, id uint32) *protogen.BookInfo {
	// Get a Book
	bookGetInput := &protogen.BookGetInput{
		Id: id,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//
	bookInfo, err := client.GetBook(ctx, bookGetInput)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("book get: %v\n", bookInfo)
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
