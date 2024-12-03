package main

import (
	"flag"
	"fmt"
	"grpc-boot-starter/config"
	"grpc-boot-starter/core/logging"
	pb "grpc-boot-starter/protogen"
	"grpc-boot-starter/services"
	"net"
	"os"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// starting the server

var (
	env = flag.String("e", "dev", "active profile, eg. dev, sit, uat, staging, prod")
	cfg = flag.String("cfg", "dev", "config file full path")
)

func getEnvOrFlagValue(flagValue, envVarName string) string {
	if flagValue == "" {
		return os.Getenv(envVarName)
	}
	return flagValue
}

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: server -cfg {path} -e {env}")
		os.Exit(1)
	}
	flag.Parse()
	//
	//
	var envName = getEnvOrFlagValue(*env, "APP_ENV")
	var cfgPath = getEnvOrFlagValue(*cfg, "APP_CFG")
	//
	fmt.Printf("running in %v, env: %v\n", cfgPath, envName)
	// load configuration
	config.LoadConf(cfgPath, envName)
	//
	logging.InitLogging()
	//
	lis, err := net.Listen("tcp", ":"+config.GetConfig().Server.Port)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to listen: %v", config.GetConfig().Server.Port)
	}
	//
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &services.HelloServiceServerImpl{})
	grpcServer.Serve(lis)
	//
}
