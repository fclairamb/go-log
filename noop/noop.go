// Package noop provides a No-Operation logger
package noop

import (
	"github.com/fclairamb/log"
)

// NewNoOpLogger creates a no-op logger
func NewNoOpLogger() log.Logger {
	return &noLogger{}
}

type noLogger struct{}

func (nl *noLogger) Debug(string, ...interface{}) {
}

func (nl *noLogger) Info(string, ...interface{}) {
}

func (nl *noLogger) Warn(string, ...interface{}) {
}

func (nl *noLogger) Error(string, ...interface{}) {
}

func (nl *noLogger) With(...interface{}) log.Logger {
	return nl
}
