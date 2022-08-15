package samples

import (
	"testing"

	"go.uber.org/zap"

	"github.com/fclairamb/go-log/logtest"
	adapter "github.com/fclairamb/go-log/zap"
)

func TestZap(t *testing.T) {
	innerLogger, _ := zap.NewProduction()
	logger := adapter.NewWrap(innerLogger.Sugar())

	logtest.TestLogger(t, logger)
}
