package services

import (
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
	healthCheck.SetServingStatus(statusSystem, healthpb.HealthCheckResponse_NOT_SERVING)
}

func UpdateServiceStatus(name string, status int) {
	healthCheck.SetServingStatus(name, healthpb.HealthCheckResponse_ServingStatus(status))
}

func SetServerServing() {
	healthCheck.SetServingStatus(statusSystem, healthpb.HealthCheckResponse_SERVING)
}

func SetServerDown() {
	healthCheck.SetServingStatus(statusSystem, healthpb.HealthCheckResponse_NOT_SERVING)
}
