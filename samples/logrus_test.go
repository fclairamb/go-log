package samples

import (
	"testing"

	"github.com/sirupsen/logrus" //nolint: depguard

	adapter "github.com/fclairamb/go-log/logrus"
	"github.com/fclairamb/go-log/logtest"
)

func TestLogrus(t *testing.T) {
	logrusLogger := logrus.New()
	logger := adapter.NewWrap(logrusLogger)

	logtest.TestLogger(t, logger)
}
