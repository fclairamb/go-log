package samples

import (
	"testing"

	"go.uber.org/zap"

	adapter "github.com/fclairamb/go-log/zap"
)

func TestZap(t *testing.T) {
	innerLogger, _ := zap.NewProduction()
	logger := adapter.NewWrap(innerLogger.Sugar())

	logger.Info("Hello world !")
}
