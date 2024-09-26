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
		Debug(msg string, args ...any)
		Info(msg string, args ...any)
		Warn(msg string, args ...any)
		Error(msg string, args ...any)
		DebugWithFields(msg string, fields Field)
		InfoWithFields(msg string, fields Field)
		WarnWithFields(msg string, fields Field)
		ErrorWithFields(msg string, fields Field)
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

func (l *logger) log(level slog.Level, msg string, args ...any) {
	l.slog.LogAttrs(context.Background(), level, msg,
		slog.String("timestamp", time.Now().Format(time.RFC3339)),
		slog.Any("data", args))
}

func (l *logger) logWithFields(level slog.Level, msg string, fields map[string]interface{}) {
	attrs := make([]slog.Attr, 0, len(fields)+1)
	attrs = append(attrs, slog.String("timestamp", time.Now().Format(time.RFC3339)))

	for k, v := range fields {
		attrs = append(attrs, slog.Any(k, v))
	}

	l.slog.LogAttrs(context.Background(), level, msg, attrs...)
}

func (l *logger) Debug(msg string, args ...any) {
	l.log(slog.LevelDebug, msg, args...)
}

func (l *logger) Info(msg string, args ...any) {
	l.log(slog.LevelInfo, msg, args...)
}

func (l *logger) Warn(msg string, args ...any) {
	l.log(slog.LevelWarn, msg, args...)
}

func (l *logger) Error(msg string, args ...any) {
	l.log(slog.LevelError, msg, args...)
}

func (l *logger) DebugWithFields(msg string, fields map[string]interface{}) {
	l.logWithFields(slog.LevelDebug, msg, fields)
}

func (l *logger) InfoWithFields(msg string, fields map[string]interface{}) {
	l.logWithFields(slog.LevelInfo, msg, fields)
}

func (l *logger) WarnWithFields(msg string, fields map[string]interface{}) {
	l.logWithFields(slog.LevelWarn, msg, fields)
}

func (l *logger) ErrorWithFields(msg string, fields map[string]interface{}) {
	l.logWithFields(slog.LevelError, msg, fields)
}
