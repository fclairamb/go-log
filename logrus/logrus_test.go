package logrus

import (
	"testing"

	"github.com/fclairamb/go-log/logtest"
)

func TestLogrus(t *testing.T) {
	logtest.TestLogger(t, New())
}
