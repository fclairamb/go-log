// Package log15 defines a simple implementation of the logging interface
package log15

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/fclairamb/go-log"
)

func convertArgs(event string, keyvals ...interface{}) []interface{} {
	args := make([]interface{}, len(keyvals)+1)
	args[0] = event
	copy(args[1:], keyvals)
	return args
}

// Debug logs key-values at debug level
func (logger *zapLogger) Debug(event string, keyvals ...interface{}) {
	logger.logger.Debug(convertArgs(event, keyvals)...)
}

// Info logs key-values at info level
func (logger *zapLogger) Info(event string, keyvals ...interface{}) {
	logger.logger.Info(convertArgs(event, keyvals)...)
}

// Warn logs key-values at warn level
func (logger *zapLogger) Warn(event string, keyvals ...interface{}) {
	logger.logger.Warn(convertArgs(event, keyvals)...)
}

// Error logs key-values at error level
func (logger *zapLogger) Error(event string, keyvals ...interface{}) {
	logger.logger.Error(convertArgs(event, keyvals)...)
}

// With adds key-values
func (logger *zapLogger) With(keyvals ...interface{}) log.Logger {
	return NewWrap(logger.logger.With(keyvals...))
}

// NewWrap creates a logger based on log15 logs
func NewWrap(logger *zap.SugaredLogger) log.Logger {
	return &zapLogger{
		logger: logger,
	}
}

// New creates a logger based on the default log15 logger
func New() (log.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("log15: cannot create logger: %w", err)
	}

	return NewWrap(logger.Sugar()), nil
}

type zapLogger struct {
	logger *zap.SugaredLogger
}
