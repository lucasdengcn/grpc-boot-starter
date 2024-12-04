package services

import (
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var (
	statusSystem = ""
	healthCheck  *health.Server
)

// RegisterService register health service to server
func RegisterHealthService(server *grpc.Server) {
	//
	healthCheck = health.NewServer()
	healthgrpc.RegisterHealthServer(server, healthCheck)
	//
	go func() {
		// asynchronously inspect dependencies and toggle serving status as needed
		next := healthpb.HealthCheckResponse_SERVING

		for {
			// toggle health
			healthCheck.SetServingStatus(statusSystem, next)
			if next == healthpb.HealthCheckResponse_SERVING {
				next = healthpb.HealthCheckResponse_NOT_SERVING
			} else {
				next = healthpb.HealthCheckResponse_SERVING
			}
			log.Info().Msgf("service status: %v", next)
			time.Sleep(5 * time.Second)
		}
	}()
}

func UpdateServiceStatus(name string, status int) {
	healthCheck.SetServingStatus(name, healthpb.HealthCheckResponse_ServingStatus(status))
}
