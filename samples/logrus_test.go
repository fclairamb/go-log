package samples

import (
	"testing"

	"github.com/sirupsen/logrus" //nolint: depguard

	adapter "github.com/fclairamb/go-log/logrus"
)

func TestLogrus(t *testing.T) {
	logrusLogger := logrus.New()
	logger := adapter.NewWrap(logrusLogger)

	logger.Info("Hello world !")
}
