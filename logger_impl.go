package go_logging

import (
    "context"
    "time"
    "golang.org/x/exp/slog"
    "os"
)

type SlogLogger struct {
    logger  *slog.Logger
    service string
}
func NewSlogLogger(service string) *SlogLogger {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    return &SlogLogger{logger: logger, service: service}
}
func (l *SlogLogger) log(level slog.Level, msg string, args ...any) {
    newArgs := append([]any{"service", l.service, "timestamp", time.Now()}, args...)
    l.logger.Log(context.Background(), level, msg, newArgs...)
}
func (l *SlogLogger) Debug(msg string, args ...any) {
    l.log(slog.LevelDebug, msg, args...)
}
func (l *SlogLogger) Info(msg string, args ...any) {
    l.log(slog.LevelInfo, msg, args...)
}
func (l *SlogLogger) Warn(msg string, args ...any) {
    l.log(slog.LevelWarn, msg, args...)
}
func (l *SlogLogger) Error(msg string, args ...any) {
    l.log(slog.LevelError, msg, args...)
}
