// Package log15 defines a simple implementation of the logging interface
package log15

import (
	"github.com/inconshreveable/log15"

	"github.com/fclairamb/go-log"
)

// Debug logs key-values at debug level
func (logger *log15Logger) Debug(event string, keyvals ...interface{}) {
	logger.logger.Debug(event, keyvals...)
}

// Info logs key-values at info level
func (logger *log15Logger) Info(event string, keyvals ...interface{}) {
	logger.logger.Info(event, keyvals...)
}

// Warn logs key-values at warn level
func (logger *log15Logger) Warn(event string, keyvals ...interface{}) {
	logger.logger.Warn(event, keyvals...)
}

// Error logs key-values at error level
func (logger *log15Logger) Error(event string, keyvals ...interface{}) {
	logger.logger.Error(event, keyvals...)
}

// With adds key-values
func (logger *log15Logger) With(keyvals ...interface{}) log.Logger {
	return NewWrap(logger.logger.New(keyvals...))
}

// NewWrap creates a logger based on log15 logs
func NewWrap(logger log15.Logger) log.Logger {
	return &log15Logger{
		logger: logger,
	}
}

// New creates a logger based on the default log15 logger
func New() log.Logger {
	return NewWrap(log15.New())
}

type log15Logger struct {
	logger log15.Logger
}
