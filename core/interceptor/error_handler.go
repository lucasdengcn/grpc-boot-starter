package interceptor

import (
	"context"
	"fmt"
	"grpc-boot-starter/core/exception"
	"grpc-boot-starter/core/logging"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleErrorGlobal recover error from panic, or handle error from normal return.
// to ensure err having value, function body MUST assign value to err at last.
func HandleErrorGlobal(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	defer func() {
		if r := recover(); r != nil {
			err0, ok := r.(error)
			if ok {
				logging.Error(ctx).Err(err0).Msgf("Recover Error: %T, %v", err0, req)
				err = exception.TranslateToGrpcStatus(err0)
			} else {
				logging.Error(ctx).Msgf("Recover Unknown Error: %T, %v", r, r)
				err = status.Errorf(codes.Unknown, fmt.Sprintf("%v", r))
			}
		}
	}()
	// Continue execution of handler after ensuring a valid token.
	resp, err0 := handler(ctx, req)
	// translate error to gRPC error
	if err0 != nil {
		logging.Error(ctx).Err(err0).Msgf("Req Error: %T, %v", err0, req)
		err = exception.TranslateToGrpcStatus(err)
	}
	return resp, err
}
