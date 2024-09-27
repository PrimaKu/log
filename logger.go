package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"time"
)

type (
	Logger interface {
		Debug(ctx context.Context, msg string, args ...any)
		Info(ctx context.Context, msg string, args ...any)
		Warn(ctx context.Context, msg string, args ...any)
		Error(ctx context.Context, msg string, args ...any)
		DebugWithFields(ctx context.Context, msg string, fields map[string]interface{})
		InfoWithFields(ctx context.Context, msg string, fields map[string]interface{})
		WarnWithFields(ctx context.Context, msg string, fields map[string]interface{})
		ErrorWithFields(ctx context.Context, msg string, fields map[string]interface{})
	}
	logger struct {
		slog        *slog.Logger
		serviceName string
	}
)

func NewLogger(serviceName string, level slog.Level, out io.Writer) Logger {
	if out == nil {
		out = os.Stdout
	}

	opts := &slog.HandlerOptions{
		Level: level,
	}

	handler := slog.NewJSONHandler(out, opts).WithAttrs([]slog.Attr{
		slog.String("service", serviceName),
	})
	slogLogger := slog.New(handler)

	return &logger{slog: slogLogger, serviceName: serviceName}
}

func (l *logger) log(ctx context.Context, level slog.Level, msg string, args ...any) {
	l.slog.LogAttrs(ctx, level, msg,
		slog.String("timestamp", time.Now().Format(time.RFC3339)),
		slog.Any("data", args))
}

func (l *logger) logWithFields(ctx context.Context, level slog.Level, msg string, fields map[string]interface{}) {
	attrs := make([]slog.Attr, 0, len(fields)+1)
	attrs = append(attrs, slog.String("timestamp", time.Now().Format(time.RFC3339)))

	for k, v := range fields {
		attrs = append(attrs, slog.Any(k, v))
	}

	l.slog.LogAttrs(ctx, level, msg, attrs...)
}

func (l *logger) Debug(ctx context.Context, msg string, args ...any) {
	l.log(ctx, slog.LevelDebug, msg, args...)
}

func (l *logger) Info(ctx context.Context, msg string, args ...any) {
	l.log(ctx, slog.LevelInfo, msg, args...)
}

func (l *logger) Warn(ctx context.Context, msg string, args ...any) {
	l.log(ctx, slog.LevelWarn, msg, args...)
}

func (l *logger) Error(ctx context.Context, msg string, args ...any) {
	l.log(ctx, slog.LevelError, msg, args...)
}

func (l *logger) DebugWithFields(ctx context.Context, msg string, fields map[string]interface{}) {
	l.logWithFields(ctx, slog.LevelDebug, msg, fields)
}

func (l *logger) InfoWithFields(ctx context.Context, msg string, fields map[string]interface{}) {
	l.logWithFields(ctx, slog.LevelInfo, msg, fields)
}

func (l *logger) WarnWithFields(ctx context.Context, msg string, fields map[string]interface{}) {
	l.logWithFields(ctx, slog.LevelWarn, msg, fields)
}

func (l *logger) ErrorWithFields(ctx context.Context, msg string, fields map[string]interface{}) {
	l.logWithFields(ctx, slog.LevelError, msg, fields)
}
