// Package gokit defines a simple implementation of the logging interface
package gokit

import (
	"fmt"
	"os"

	gklog "github.com/go-kit/log"
	gklevel "github.com/go-kit/log/level"

	"github.com/fclairamb/go-log"
)

func (logger *gKLogger) checkError(err error) {
	if err != nil {
		fmt.Println("Logging faced this error: ", err) //nolint:forbidigo
	}
}

func (logger *gKLogger) log(gklogger gklog.Logger, event string, keyvals ...interface{}) {
	keyvals = append(keyvals, "event", event)
	logger.checkError(gklogger.Log(keyvals...))
}

// Debug logs key-values at debug level
func (logger *gKLogger) Debug(event string, keyvals ...interface{}) {
	logger.log(gklevel.Debug(logger.logger), event, keyvals...)
}

// Info logs key-values at info level
func (logger *gKLogger) Info(event string, keyvals ...interface{}) {
	logger.log(gklevel.Info(logger.logger), event, keyvals...)
}

// Warn logs key-values at warn level
func (logger *gKLogger) Warn(event string, keyvals ...interface{}) {
	logger.log(gklevel.Warn(logger.logger), event, keyvals...)
}

// Error logs key-values at error level
func (logger *gKLogger) Error(event string, keyvals ...interface{}) {
	logger.log(gklevel.Error(logger.logger), event, keyvals...)
}

func (logger *gKLogger) Panic(event string, keyvals ...interface{}) {
	logger.Error(event, keyvals...)

	panic(fmt.Errorf("%s: %s", event, keyvals)) //nolint:err113
}

// With adds key-values
func (logger *gKLogger) With(keyvals ...interface{}) log.Logger {
	return NewWrap(gklog.With(logger.logger, keyvals...))
}

// NewWrap creates a logger based on go-kit logs
func NewWrap(logger gklog.Logger) log.Logger {
	return &gKLogger{
		logger: logger,
	}
}

// New creates a logger based on go-kit logs but with some default parameters
func New() log.Logger {
	return NewWrap(gklog.NewLogfmtLogger(gklog.NewSyncWriter(os.Stdout)))
}

type gKLogger struct {
	logger gklog.Logger
}

const callDepth = 5

var (
	// GKDefaultCaller adds a "caller" property
	GKDefaultCaller = gklog.Caller(callDepth)
	// GKDefaultTimestampUTC adds a "ts" property
	GKDefaultTimestampUTC = gklog.DefaultTimestampUTC
)
