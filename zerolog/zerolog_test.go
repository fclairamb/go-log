package zerolog

import (
	"testing"

	"github.com/fclairamb/go-log/logtest"
)

func TestZerolog(t *testing.T) {
	logtest.TestLogger(t, New())
}
