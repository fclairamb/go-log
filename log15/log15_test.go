package log15

import (
	"testing"

	"github.com/fclairamb/go-log/logtest"
)

func TestLog15(t *testing.T) {
	logtest.TestLogger(t, New())
}
