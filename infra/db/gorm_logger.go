package db

import (
	"context"
	"fmt"
	"grpc-boot-starter/core/logging"
	"time"

	"github.com/rs/zerolog"

	"gorm.io/gorm/logger"
)

type GormLogger struct {
}

func (l *GormLogger) LogMode(logger.LogLevel) logger.Interface {
	return l
}

func (l *GormLogger) Error(ctx context.Context, msg string, opts ...interface{}) {
	fmt.Println("Error GormLogger")
	logging.Error(ctx).Msg(fmt.Sprintf(msg, opts...))
}

func (l *GormLogger) Warn(ctx context.Context, msg string, opts ...interface{}) {
	fmt.Println("Warn GormLogger")
	logging.Warn(ctx).Msg(fmt.Sprintf(msg, opts...))
}

func (l *GormLogger) Info(ctx context.Context, msg string, opts ...interface{}) {
	fmt.Println("Info GormLogger")
	logging.Info(ctx).Msg(fmt.Sprintf(msg, opts...))
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, f func() (string, int64), err error) {
	var event *zerolog.Event

	if err != nil {
		event = logging.Error(ctx)
	} else {
		event = logging.Info(ctx)
	}

	event.Str("duration", fmt.Sprintf("%v", time.Since(begin)))

	sql, rows := f()
	if sql != "" {
		event.Str("sql", sql)
	}
	if rows > -1 {
		event.Int64("rows", rows)
	}

	event.Send()

	return
}
