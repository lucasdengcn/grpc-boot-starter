package interceptor

import (
	"context"
	"grpc-boot-starter/core/models"
	"strings"
	"time"

	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// CtxLogger attach logger to context, put tracing, spanId, correlationId into logger context.
func CtxLogger(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// get zerolog
	z := log.Logger
	// return if zerolog is disabled
	if z.GetLevel() == zerolog.Disabled {
		return handler(ctx, req)
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	ctx, l := buildContextLogger(ctx, md, &z)
	l.Info().Msgf("Income req: %T, %v", req, req)
	// Continue execution of handler after ensuring a valid token.
	begin := time.Now()
	resp, err := handler(ctx, req)
	duration := time.Since(begin)
	l.Info().Msgf("Income req: %T, %v, duration: %v", req, req, duration)
	return resp, err
}

func mdValue(md metadata.MD, key string) string {
	vals := md.Get(key)
	if len(vals) == 0 {
		return ""
	}
	return vals[0]
}

func buildContextLogger(ctx context.Context, md metadata.MD, z *zerolog.Logger) (context.Context, *zerolog.Logger) {
	l := z.With().Logger()
	traceparent := mdValue(md, "traceparent")
	traceID := ""
	spanID := ""
	correlationID := ""
	corrCtx := &models.CorrelationCtx{}
	if traceparent != "" {
		parts := strings.Split(traceparent, "-")
		if len(parts) == 4 {
			traceID = parts[1]
			spanID = parts[2]
			corrCtx.TraceId = traceID
			corrCtx.SpanId = spanID
		}
	}
	//
	if traceID == "" || spanID == "" {
		correlationID = xid.New().String()
		corrCtx.Id = correlationID
	}
	//
	l.UpdateContext(func(c zerolog.Context) zerolog.Context {
		zc := c
		if correlationID != "" {
			zc = zc.Str(models.CorrelationIdFieldName, correlationID)
		}
		if traceID != "" {
			zc = zc.Str(models.TraceIdFieldName, traceID)
		}
		if spanID != "" {
			zc = zc.Str(models.SpanIdFieldName, spanID)
		}
		return zc
	})
	//
	ctx = context.WithValue(ctx, models.CorrelationCtxKey, corrCtx)
	return l.WithContext(ctx), &l
}
