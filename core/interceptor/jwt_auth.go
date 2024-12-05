package interceptor

import (
	"context"
	"grpc-boot-starter/core/logging"
	"grpc-boot-starter/core/models"
	"grpc-boot-starter/core/security"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ensureValidToken ensures a valid token exists within a request's metadata. If
// the token is missing or invalid, the interceptor blocks execution of the
// handler and returns an error. Otherwise, the interceptor invokes the unary
// handler.
func EnsureValidToken(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	logging.Info(ctx).Msgf("Income req: %T, %v", req, req)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, models.NewValidationError(ctx, "MD_404", "missing metadata")
	}
	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	newCtx, err := valid(ctx, md["authorization"])
	if err != nil {
		return nil, err
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(newCtx, req)
}

// valid validates the authorization.
func valid(ctx context.Context, authorization []string) (context.Context, error) {
	if len(authorization) < 1 {
		return ctx, models.NewAuthError(ctx, "AUTH_404", "auth token required")
	}
	tokenString := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	token, err := jwt.ParseWithClaims(tokenString, &security.AuthClaims{}, security.PublicJwtKeyfuncCtx(ctx))
	if err != nil || !token.Valid {
		logging.Error(ctx).Err(err).Msgf("Token Invalid:%s", tokenString)
		return ctx, models.NewAuthError(ctx, "AUTH_401", "invalid auth token")
	}
	// save current user
	return security.SaveCurrentUser(ctx, token), nil
}
