package main

import (
	"flag"
	"fmt"
	"grpc-boot-starter/config"
	"grpc-boot-starter/core/logging"
	"grpc-boot-starter/infra/db"
	"grpc-boot-starter/migration"
	"grpc-boot-starter/services"
	"net"
	"os"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
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
	//
	grpcServer := grpc.NewServer()
	//
	services.RegisterHealthService(grpcServer)
	services.RegisterHelloService(grpcServer)
	//
	grpcServer.Serve(lis)
	//
}
