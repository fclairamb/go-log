// Package testing provides an adapter for the go test testing logger
package testing

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fclairamb/go-log"
	"github.com/fclairamb/go-log/level"
)

// NewTestLogger creates a logger based on the test framework
func NewTestLogger(t *testing.T, level level.Level) log.Logger {
	return &testLogger{
		t:       t,
		context: nil,
		level:   level,
	}
}

// testLogger is a logger based on the test framework
type testLogger struct {
	t       *testing.T
	context []interface{}
	level   level.Level
}

func (t *testLogger) log(level level.Level, msg string, args []interface{}) {
	t.t.Helper()

	if !t.level.ShouldLog(level) {
		return
	}

	sb := strings.Builder{}
	sb.WriteString(level.String())
	sb.WriteString(": ")
	sb.WriteString(msg)

	if t.context != nil {
		args = append(args, t.context...)
	}

	if len(args) > 0 {
		sb.WriteString(" / ")

		for i := 0; i < len(args); i++ {
			if i%2 == 0 {
				if i > 0 {
					sb.WriteString(", ")
				}

				sb.WriteString(fmt.Sprintf("%v", args[i]))
			} else {
				sb.WriteString(fmt.Sprintf("=\"%v\"", args[i]))
			}
		}
	}

	t.t.Log(sb.String())
}

func (t *testLogger) Debug(msg string, args ...interface{}) {
	t.t.Helper()
	t.log(level.Debug, msg, args)
}

func (t *testLogger) Info(msg string, args ...interface{}) {
	t.t.Helper()
	t.log(level.Info, msg, args)
}

func (t *testLogger) Warn(msg string, args ...interface{}) {
	t.t.Helper()
	t.log(level.Warning, msg, args)
}

func (t *testLogger) Error(msg string, args ...interface{}) {
	t.t.Helper()
	t.log(level.Error, msg, args)
}

// We don't _log_ anything but we need to keep a consistent behavior

func (t *testLogger) Panic(msg string, args ...interface{}) {
	panic(fmt.Errorf("%s: %s", msg, args)) //nolint:goerr113
}

func (t *testLogger) With(args ...interface{}) log.Logger {
	logger := &testLogger{
		t:       t.t,
		context: nil,
		level:   t.level,
	}

	if t.context != nil {
		logger.context = append(logger.context, t.context...)
	}

	if len(args) > 0 {
		logger.context = append(logger.context, args...)
	}

	return logger
}
