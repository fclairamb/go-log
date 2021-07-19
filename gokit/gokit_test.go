package gokit

import (
	"os"
	"testing"

	"github.com/fclairamb/go-log/logtest"
)

func TestGoKit(t *testing.T) {
	logtest.TestLogger(t, NewGKLoggerStdout())
}

type badLogger struct{}

func (badLogger) Log(keyvals ...interface{}) error {
	return os.ErrInvalid
}

func TestError(t *testing.T) {
	logger := NewGKLogger(badLogger{})
	logger.Info("test")
}
