// Package log15 defines a simple implementation of the logging interface
package log15

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/fclairamb/go-log"
)

// Debug logs key-values at debug level
func (logger *zapLogger) Debug(event string, keyvals ...interface{}) {
	logger.logger.Debugw(event, keyvals...)
}

// Info logs key-values at info level
func (logger *zapLogger) Info(event string, keyvals ...interface{}) {
	logger.logger.Infow(event, keyvals...)
}

// Warn logs key-values at warn level
func (logger *zapLogger) Warn(event string, keyvals ...interface{}) {
	logger.logger.Warnw(event, keyvals...)
}

// Error logs key-values at error level
func (logger *zapLogger) Error(event string, keyvals ...interface{}) {
	logger.logger.Errorw(event, keyvals...)
}

func (logger *zapLogger) Panic(event string, keyvals ...interface{}) {
	logger.logger.Panicw(event, keyvals...)
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
func New(options ...zap.Option) (log.Logger, error) {
	logger, err := zap.NewProduction(options...)
	if err != nil {
		return nil, fmt.Errorf("zap: cannot create logger: %w", err)
	}

	return NewWrap(logger.Sugar()), nil
}

type zapLogger struct {
	logger *zap.SugaredLogger
}
