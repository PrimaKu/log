package logger

import (
	"log/slog"
	"strings"
)

func LogLevelFromStr(logLevelStr string) slog.Level {
	switch {
	case strings.EqualFold(logLevelStr, "DEBUG"):
		return slog.LevelDebug
	case strings.EqualFold(logLevelStr, "INFO"):
		return slog.LevelInfo
	case strings.EqualFold(logLevelStr, "WARN"):
		return slog.LevelWarn
	default:
		return slog.LevelError
	}
}
