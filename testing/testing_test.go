package testing

import (
	gotest "testing"

	"github.com/fclairamb/go-log/level"
	"github.com/fclairamb/go-log/logtest"
)

func TestTesting(t *gotest.T) {
	logtest.TestLogger(t, NewTestLogger(t, level.Debug))
}

func TestLevel(t *gotest.T) {
	NewTestLogger(t, level.Info).Debug("Should not be displayed")
}
