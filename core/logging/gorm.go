package logging

import (
	"context"
	"fmt"
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
	Error(ctx).Msg(fmt.Sprintf(msg, opts...))
}

func (l *GormLogger) Warn(ctx context.Context, msg string, opts ...interface{}) {
	Warn(ctx).Msg(fmt.Sprintf(msg, opts...))
}

func (l *GormLogger) Info(ctx context.Context, msg string, opts ...interface{}) {
	Info(ctx).Msg(fmt.Sprintf(msg, opts...))
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, f func() (string, int64), err error) {
	var event *zerolog.Event

	if err != nil {
		event = Debug(ctx)
	} else {
		event = Trace(ctx)
	}

	var dur_key string

	switch zerolog.DurationFieldUnit {
	case time.Nanosecond:
		dur_key = "elapsed_ns"
	case time.Microsecond:
		dur_key = "elapsed_us"
	case time.Millisecond:
		dur_key = "elapsed_ms"
	case time.Second:
		dur_key = "elapsed"
	case time.Minute:
		dur_key = "elapsed_min"
	case time.Hour:
		dur_key = "elapsed_hr"
	default:
		Error(ctx).Interface("zerolog.DurationFieldUnit", zerolog.DurationFieldUnit).Msg("unknown value for DurationFieldUnit")
		dur_key = "elapsed_"
	}

	event.Dur(dur_key, time.Since(begin))

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
