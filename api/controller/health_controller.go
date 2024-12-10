package controller

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

// RegisterHealthCheck register health service to server
func RegisterHealthCheck(server *grpc.Server) {
	//
	healthCheck = health.NewServer()
	healthgrpc.RegisterHealthServer(server, healthCheck)
	healthCheck.SetServingStatus(statusSystem, healthpb.HealthCheckResponse_NOT_SERVING)
}

func UpdateStatus(name string, status int) {
	healthCheck.SetServingStatus(name, healthpb.HealthCheckResponse_ServingStatus(status))
}

func UpdateServerServing() {
	healthCheck.SetServingStatus(statusSystem, healthpb.HealthCheckResponse_SERVING)
}

func UpdateServerDown() {
	healthCheck.SetServingStatus(statusSystem, healthpb.HealthCheckResponse_NOT_SERVING)
}
