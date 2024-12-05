package interceptor

import (
	"context"
	"errors"
	"grpc-boot-starter/core/logging"
	"grpc-boot-starter/core/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleErrorGlobal(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// Continue execution of handler after ensuring a valid token.
	resp, err := handler(ctx, req)
	if err != nil {
		logging.Error(ctx).Err(err).Msgf("Req Error: %T, %v, %v", err, req, errors.Is(err, &models.AuthError{}))
		// translate error to gRPC error
		switch {
		case errors.Is(err, &models.RepositoryError{}):
			return resp, status.Errorf(codes.Internal, err.Error())
		case errors.Is(err, &models.AuthError{}):
			return resp, status.Errorf(codes.Unauthenticated, err.Error())
		case errors.Is(err, &models.ACLError{}):
			return resp, status.Errorf(codes.PermissionDenied, err.Error())
		case errors.Is(err, &models.EntityNotFoundError{}):
			return resp, status.Errorf(codes.NotFound, err.Error())
		case errors.Is(err, &models.ServiceError{}):
			return resp, status.Errorf(codes.Internal, err.Error())
		case errors.Is(err, &models.ValidationError{}):
			return resp, status.Errorf(codes.InvalidArgument, err.Error())
		default:
			return resp, status.Errorf(codes.Unknown, err.Error())
		}
	}
	return resp, err
}
