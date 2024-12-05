package server

import (
	"crypto/tls"
	"grpc-boot-starter/core/config"
	"grpc-boot-starter/core/interceptor"
	"grpc-boot-starter/protogen"
	"grpc-boot-starter/services"
	"net"
	"path/filepath"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/stats/opentelemetry"
)

var kaep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

type AppServer struct {
	serv *grpc.Server
	lis  *net.Listener
}

func NewAppServer(otelOpts opentelemetry.Options) *AppServer {
	// mTLS
	basePath := config.GetConfig().Application.WorkingPath
	cert, err := tls.LoadX509KeyPair(filepath.Join(basePath, "secrets/x509/server_cert.pem"), filepath.Join(basePath, "secrets/x509/server_key.pem"))
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to load key pair: %s", err)
	}
	// opts
	opts := []grpc.ServerOption{
		opentelemetry.ServerOption(otelOpts),
		// The following grpc.ServerOption adds an interceptor for all unary
		// RPCs. To configure an interceptor for streaming RPCs, see:
		// https://godoc.org/google.golang.org/grpc#StreamInterceptor
		grpc.ChainUnaryInterceptor(
			interceptor.CtxLogger,
			interceptor.HandleErrorGlobal,
			interceptor.EnsureValidToken,
		),
		// Enable TLS for all incoming connections.
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		// KeepAlive settings
		grpc.KeepaliveEnforcementPolicy(kaep),
		grpc.KeepaliveParams(kasp),
	}
	//
	grpcServer := grpc.NewServer(opts...)
	//
	return &AppServer{serv: grpcServer}
}

// register services to gRPC server
func (s *AppServer) registerServiceServers() {
	//
	service.RegisterChannelzServiceToServer(s.serv)
	// hook services
	services.RegisterHealthService(s.serv)
	protogen.RegisterBookServiceServer(s.serv, InitializeBookService())
	protogen.RegisterHelloServiceServer(s.serv, InitializeHelloService())
	// update service status
	services.SetServerServing()
	//
	log.Info().Msg("Server services init done.")
}

func (s *AppServer) Start() error {
	// have service ready first.
	s.registerServiceServers()
	// Start listening on port.
	port := config.GetConfig().Server.Port
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Error().Err(err).Msgf("failed to listen: %v", port)
		return err
	}
	log.Info().Msgf("Server start and listen at: %v", port)
	s.serv.Serve(lis)
	return nil
}

func (s *AppServer) Stop() {
	// Stop the server.
	s.serv.GracefulStop()
}
