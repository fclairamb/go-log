package gokit

import (
	"testing"

	"github.com/fclairamb/log/logtest"
)

func TestGoKit(t *testing.T) {
	logtest.TestLogger(t, NewGKLoggerStdout())
}
