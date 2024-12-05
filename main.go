package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"grpc-boot-starter/core/config"
	"grpc-boot-starter/core/interceptor"
	"grpc-boot-starter/core/logging"
	"grpc-boot-starter/infra/db"
	"grpc-boot-starter/migration"
	"grpc-boot-starter/protogen"
	"grpc-boot-starter/server"
	"grpc-boot-starter/services"
	"net"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// starting the server

var (
	env = flag.String("e", "dev", "active profile, eg. dev, sit, uat, staging, prod")
	cfg = flag.String("w", "dev", "working directory full path")
)

func getEnvOrFlagValue(flagValue, envVarName string) string {
	if flagValue == "" {
		return os.Getenv(envVarName)
	}
	return flagValue
}

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: server -w {path} -e {env}")
		os.Exit(1)
	}
	flag.Parse()
	//
	//
	var envName = getEnvOrFlagValue(*env, "APP_ENV")
	var workingPath = getEnvOrFlagValue(*cfg, "APP_BASE")
	//
	fmt.Printf("running in %v, env: %v\n", workingPath, envName)
	// load configuration
	config.LoadConf(workingPath, envName)
	//
	logging.InitLogging()
	//
	db.ConnectDB()
	//
	migration.Migrate()
	//
	lis, err := net.Listen("tcp", ":"+config.GetConfig().Server.Port)
	if err != nil {
		db.Close()
		log.Fatal().Err(err).Msgf("failed to listen: %v", config.GetConfig().Server.Port)
	}
	log.Info().Msgf("Listen at port: %v", config.GetConfig().Server.Port)
	//
	grpcServer := setupGrpcServer()
	registerServiceServers(grpcServer)
	grpcServer.Serve(lis)
}

func setupGrpcServer() *grpc.Server {
	// mTLS
	basePath := config.GetConfig().Application.WorkingPath
	cert, err := tls.LoadX509KeyPair(filepath.Join(basePath, "secrets/x509/server_cert.pem"), filepath.Join(basePath, "secrets/x509/server_key.pem"))
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to load key pair: %s", err)
	}
	// opts
	opts := []grpc.ServerOption{
		// The following grpc.ServerOption adds an interceptor for all unary
		// RPCs. To configure an interceptor for streaming RPCs, see:
		// https://godoc.org/google.golang.org/grpc#StreamInterceptor
		grpc.ChainUnaryInterceptor(
			interceptor.CtxLogger,
			interceptor.EnsureValidToken,
		),
		// Enable TLS for all incoming connections.
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	//
	grpcServer := grpc.NewServer(opts...)
	return grpcServer
}

func registerServiceServers(grpcServer *grpc.Server) {
	// hook services
	services.RegisterHealthService(grpcServer)
	protogen.RegisterBookServiceServer(grpcServer, server.InitializeBookService())
	protogen.RegisterHelloServiceServer(grpcServer, server.InitializeHelloService())
	services.SetServerServing()
}
