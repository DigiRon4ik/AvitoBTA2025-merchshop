// Package logger need it for the work of Logger.
package logger

import (
	"log/slog"
	"os"
)

// Config - configuration for Logger.
type Config struct {
	Level     slog.Level `envconfig:"LEVEL" default:"info"`
	AddSource bool       `envconfig:"ADDSOURCE" default:"false"`
}

// Init - Initialize the logger.
func Init(cfg Config) *slog.Logger {
	logLevel := &slog.LevelVar{} // INFO
	logLevel.Set(cfg.Level)

	opts := &slog.HandlerOptions{
		AddSource: cfg.AddSource,
		Level:     logLevel,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)

	logger := slog.New(handler)

	slog.SetDefault(logger)

	return logger
}
