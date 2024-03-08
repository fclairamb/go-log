package slog

import (
	"testing"

	"github.com/fclairamb/go-log/logtest"
)

func TestSLog(t *testing.T) {
	logtest.TestLogger(t, New())
}
