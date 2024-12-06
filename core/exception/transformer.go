package exception

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TranslateToGrpcStatus(err error) error {
	if err != nil {
		switch {
		case errors.Is(err, &RepositoryError{}):
			err = status.Errorf(codes.Internal, err.Error())
		case errors.Is(err, &AuthError{}):
			err = status.Errorf(codes.Unauthenticated, err.Error())
		case errors.Is(err, &ACLError{}):
			err = status.Errorf(codes.PermissionDenied, err.Error())
		case errors.Is(err, &EntityNotFoundError{}):
			err = status.Errorf(codes.NotFound, err.Error())
		case errors.Is(err, &ServiceError{}):
			err = status.Errorf(codes.Internal, err.Error())
		case errors.Is(err, &ValidationError{}):
			err = status.Errorf(codes.InvalidArgument, err.Error())
		default:
			err = status.Errorf(codes.Unknown, err.Error())
		}
	}
	return err
}
