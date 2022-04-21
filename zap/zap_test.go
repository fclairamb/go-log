package log15

import (
	"testing"

	"github.com/fclairamb/go-log/logtest"
)

func TestLog15(t *testing.T) {
	logger, err := New()
	if err != nil {
		t.Fatal(err)
	}

	logtest.TestLogger(t, logger)
}
