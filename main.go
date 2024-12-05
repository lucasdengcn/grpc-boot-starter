package main

import (
	"context"
	"flag"
	"fmt"
	"grpc-boot-starter/core/config"
	"grpc-boot-starter/core/logging"
	"grpc-boot-starter/core/otel"
	"grpc-boot-starter/infra/db"
	"grpc-boot-starter/migration"
	"grpc-boot-starter/server"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
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
	var envName = getEnvOrFlagValue(*env, "APP_ENV")
	var workingPath = getEnvOrFlagValue(*cfg, "APP_BASE")
	//
	fmt.Printf("running in %v, env: %v\n", workingPath, envName)
	// load configuration
	config.LoadConf(workingPath, envName)
	// init global logger
	logging.InitLogging()
	// prepare db connection pool
	db.ConnectDB()
	// migrate db schemas
	migration.Migrate()
	// init OTEL tracing
	opts, err := otel.InitProviders(context.Background())
	if err != nil {
		cleanUp()
		os.Exit(1)
	}
	//
	appServer := server.NewAppServer(opts)
	//
	go func() {
		err := appServer.Start()
		if err != nil {
			cleanUp()
			os.Exit(1)
		}
	}()
	shutdownHook(appServer)
}

func shutdownHook(appServer *server.AppServer) {
	// quit channel to receive stop
	quit := make(chan os.Signal, 1)
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT, Ctrl+C
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	//
	<-quit
	//
	log.Info().Msg("Shutdown Server within 5 seconds ...")
	appServer.Stop()
	//
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cleanUp()
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Info().Msg("timeout of 5 seconds.")
	}
	log.Info().Msg("Server exiting")
}

func cleanUp() {
	db.Close()
	otel.Shutdown(context.Background())
}
