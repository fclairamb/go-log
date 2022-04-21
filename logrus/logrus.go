// Package logrus provides a logger based on logrus
package logrus

import (
	"fmt"

	"github.com/sirupsen/logrus" //nolint: depguard

	"github.com/fclairamb/go-log"
)

func keyValsToFields(keyvals ...interface{}) logrus.Fields {
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "*** missing value ***")
	}

	fields := make(logrus.Fields)

	for i := 0; i < len(keyvals); i += 2 {
		fields[fmt.Sprint(keyvals[i])] = keyvals[i+1]
	}

	return fields
}

// Debug logs key-values at debug level
func (logger *logrusLogger) Debug(event string, keyvals ...interface{}) {
	logger.logger.WithFields(keyValsToFields(keyvals...)).Debug(event)
}

// Info logs key-values at info level
func (logger *logrusLogger) Info(event string, keyvals ...interface{}) {
	logger.logger.WithFields(keyValsToFields(keyvals...)).Info(event)
}

// Warn logs key-values at warn level
func (logger *logrusLogger) Warn(event string, keyvals ...interface{}) {
	logger.logger.WithFields(keyValsToFields(keyvals...)).Warn(event)
}

// Error logs key-values at error level
func (logger *logrusLogger) Error(event string, keyvals ...interface{}) {
	logger.logger.WithFields(keyValsToFields(keyvals...)).Error(event)
}

func (logger *logrusLogger) Panic(event string, keyvals ...interface{}) {
	logger.logger.WithFields(keyValsToFields(keyvals...)).Panic(event)
}

// With adds key-values
func (logger *logrusLogger) With(keyvals ...interface{}) log.Logger {
	return &logrusLogger{
		logger: logger.logger.WithFields(keyValsToFields(keyvals...)),
	}
}

// NewWrap creates a logger based on logrus logs
func NewWrap(logger *logrus.Logger) log.Logger {
	return &logrusLogger{
		logger: logger,
	}
}

// New creates a logger based on the default logrus logger
func New() log.Logger {
	return NewWrap(logrus.New())
}

type logrusLogger struct {
	logger logrus.FieldLogger
}
