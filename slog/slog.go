// Package slog defines a simple implementation of the logging interface
package slog

import (
	"context"
	"fmt"

	stdSLog "log/slog"

	"github.com/fclairamb/go-log"
)

func (logger *sLogger) log(level stdSLog.Level, event string, keyvals ...interface{}) {
	logger.logger.Log(context.Background(), level, event, keyvals...)
}

// Debug logs key-values at debug level
func (logger *sLogger) Debug(event string, keyvals ...interface{}) {
	logger.log(stdSLog.LevelDebug, event, keyvals...)
}

// Info logs key-values at info level
func (logger *sLogger) Info(event string, keyvals ...interface{}) {
	logger.log(stdSLog.LevelInfo, event, keyvals...)
}

// Warn logs key-values at warn level
func (logger *sLogger) Warn(event string, keyvals ...interface{}) {
	logger.log(stdSLog.LevelWarn, event, keyvals...)
}

// Error logs key-values at error level
func (logger *sLogger) Error(event string, keyvals ...interface{}) {
	logger.log(stdSLog.LevelError, event, keyvals...)
}

func (logger *sLogger) Panic(event string, keyvals ...interface{}) {
	logger.Error(event, keyvals...)

	panic(fmt.Errorf("%s: %s", event, keyvals)) //nolint:err113
}

// With adds key-values
func (logger *sLogger) With(keyvals ...interface{}) log.Logger {
	return NewWrap(logger.logger.With(keyvals...))
}

// NewWrap creates a logger based on go-kit logs
func NewWrap(logger *stdSLog.Logger) log.Logger {
	return &sLogger{
		logger: logger,
	}
}

// New creates a logger based on go-kit logs but with some default parameters
func New() log.Logger {
	return NewWrap(stdSLog.Default())
}

type sLogger struct {
	logger *stdSLog.Logger
}
