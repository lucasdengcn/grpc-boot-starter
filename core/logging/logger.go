package logging

import (
	"context"
	"grpc-boot-starter/config"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const loggerKey = "_zero_logger_"

func InitLogging() {
	cfg := config.GetConfig().Logging
	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		log.Info().Msgf("configuration logging.level: invalid. %v", cfg.Level)
		level = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(level)
	// logger
	var writer io.Writer
	if cfg.Format == "json" {
		writer = os.Stdout
	} else {
		writer = zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	}
	log.Logger = zerolog.New(writer).With().Timestamp().Caller().Logger()
}

func With(ctx context.Context) *zerolog.Logger {
	return log.Ctx(ctx)
}

// Get context logger
func Get(ctx context.Context) *zerolog.Logger {
	l := ctx.Value(loggerKey)
	if l == nil {
		return &log.Logger
	}
	ll := l.(zerolog.Logger)
	return &ll
}

// Trace starts a new message with trace level.
func Trace(ctx context.Context) *zerolog.Event {
	if ctx == nil {
		return log.Trace()
	}
	return Get(ctx).Trace()
}

// Debug starts a new message with debug level.
func Debug(ctx context.Context) *zerolog.Event {
	if ctx == nil {
		return log.Debug()
	}
	return Get(ctx).Debug()
}

// Info starts a new message with info level.
func Info(ctx context.Context) *zerolog.Event {
	if ctx == nil {
		return log.Info()
	}
	return Get(ctx).Info()
}

// Warn starts a new message with warn level.
func Warn(ctx context.Context) *zerolog.Event {
	if ctx == nil {
		return log.Warn()
	}
	return Get(ctx).Warn()
}

// Error starts a new message with error level.
func Error(ctx context.Context) *zerolog.Event {
	if ctx == nil {
		return log.Error()
	}
	return Get(ctx).Error()
}

// Fatal starts a new message with fatal level.
func Fatal(ctx context.Context) *zerolog.Event {
	if ctx == nil {
		return log.Fatal()
	}
	return Get(ctx).Fatal()
}

// Panic starts a new message with panic level.
func Panic(ctx context.Context) *zerolog.Event {
	if ctx == nil {
		return log.Panic()
	}
	return Get(ctx).Panic()
}
