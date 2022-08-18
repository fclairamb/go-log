// Package zerolog defines a simple implementation of the logging interface
package zerolog

import (
	"fmt"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"

	"github.com/fclairamb/go-log"
)

func addLog(event *zerolog.Event, msg string, keyvals ...interface{}) {
	for i := 0; i < len(keyvals)-1; i += 2 {
		event = event.Interface(fmt.Sprint(keyvals[i]), keyvals[i+1])
	}

	event.Msg(msg)
}

//nolint:gocritic
func addContext(context zerolog.Context, keyvals ...interface{}) zerolog.Context {
	for i := 0; i < len(keyvals)-1; i += 2 {
		context = context.Interface(fmt.Sprint(keyvals[i]), keyvals[i+1])
	}

	return context
}

// Debug logs key-values at debug level
func (logger *zeroLogLogger) Debug(event string, keyvals ...interface{}) {
	addLog(logger.logger.Debug(), event, keyvals...)
}

// Info logs key-values at info level
func (logger *zeroLogLogger) Info(event string, keyvals ...interface{}) {
	addLog(logger.logger.Info(), event, keyvals...)
}

// Warn logs key-values at warn level
func (logger *zeroLogLogger) Warn(event string, keyvals ...interface{}) {
	addLog(logger.logger.Warn(), event, keyvals...)
}

// Error logs key-values at error level
func (logger *zeroLogLogger) Error(event string, keyvals ...interface{}) {
	addLog(logger.logger.Error(), event, keyvals...)
}

func (logger *zeroLogLogger) Panic(event string, keyvals ...interface{}) {
	addLog(logger.logger.Panic(), event, keyvals...)
}

// With adds key-values
func (logger *zeroLogLogger) With(keyvals ...interface{}) log.Logger {
	lgr := addContext(logger.logger.With(), keyvals...).Logger()

	return NewWrap(&lgr)
}

// NewWrap creates a logger based on log15 logs
func NewWrap(logger *zerolog.Logger) log.Logger {
	return &zeroLogLogger{
		logger: logger,
	}
}

// New creates a logger based on the default log15 logger
func New() log.Logger {
	return NewWrap(&zlog.Logger)
}

type zeroLogLogger struct {
	logger *zerolog.Logger
}
