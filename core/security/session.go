package security

import (
	"context"
	"grpc-boot-starter/core/logging"

	jwt "github.com/golang-jwt/jwt/v5"
)

// SaveCurrentUser to attache current user to request context
func SaveCurrentUser(ctx context.Context, token *jwt.Token) context.Context {
	principle := NewPrinciple(token)
	//c.Set(PrincipleContextKey, principle)
	logging.Debug(ctx).Msgf("Current user is: %s", principle)
	return context.WithValue(ctx, PrincipleContextKey, principle)
}

// CurrentUser return current user attach to request context
func CurrentUser(ctx context.Context) *Principle {
	val := ctx.Value(PrincipleContextKey)
	if val == nil {
		return nil
	}
	return val.(*Principle)
}

// CurrentUser return current user attach to request context
func IsAuthenticated(ctx context.Context) bool {
	val := ctx.Value(PrincipleContextKey)
	if val == nil {
		return false
	}
	return true
}
