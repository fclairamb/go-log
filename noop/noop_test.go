package noop

import (
	"testing"

	"github.com/fclairamb/log/logtest"
)

func TestNoOp(t *testing.T) {
	logtest.TestLogger(t, NewNoOpLogger())
}
