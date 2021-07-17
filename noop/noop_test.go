package noop

import (
	"testing"

	"github.com/fclairamb/go-log/logtest"
)

func TestNoOp(t *testing.T) {
	logtest.TestLogger(t, NewNoOpLogger())
}
