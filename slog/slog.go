// Package gokit defines a simple implementation of the logging interface
package slog

import (
	"context"
	"fmt"

	stdLog "log/slog"

	"github.com/fclairamb/go-log"
)

func (logger *sLogger) checkError(err error) {
	if err != nil {
		fmt.Println("Logging faced this error: ", err)
	}
}

func (logger *sLogger) log(level stdLog.Level, event string, keyvals ...interface{}) {
	logger.logger.Log(logger.ctx, level, event, keyvals...)
}

// Debug logs key-values at debug level
func (logger *sLogger) Debug(event string, keyvals ...interface{}) {
	logger.log(stdLog.LevelDebug, event, keyvals...)
}

// Info logs key-values at info level
func (logger *sLogger) Info(event string, keyvals ...interface{}) {
	logger.log(stdLog.LevelInfo, event, keyvals...)
}

// Warn logs key-values at warn level
func (logger *sLogger) Warn(event string, keyvals ...interface{}) {
	logger.log(stdLog.LevelWarn, event, keyvals...)
}

// Error logs key-values at error level
func (logger *sLogger) Error(event string, keyvals ...interface{}) {
	logger.log(stdLog.LevelError, event, keyvals...)
}

func (logger *sLogger) Panic(event string, keyvals ...interface{}) {
	logger.Error(event, keyvals...)

	panic(fmt.Errorf("%s: %s", event, keyvals)) //nolint:goerr113
}

// With adds key-values
func (logger *sLogger) With(keyvals ...interface{}) log.Logger {
	//return NewWrap(slog.With(logger.logger, keyvals...))
	return NewWrap(logger.logger.With(keyvals...))
}

// NewWrap creates a logger based on go-kit logs
func NewWrap(logger *stdLog.Logger) log.Logger {
	return &sLogger{
		logger: logger,
	}
}

// New creates a logger based on go-kit logs but with some default parameters
func New() log.Logger {
	return NewWrap(stdLog.Default())
}

type sLogger struct {
	logger *stdLog.Logger
	ctx    context.Context
}
