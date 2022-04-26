package samples

import (
	"os"
	"testing"

	gklog "github.com/go-kit/log"

	adapter "github.com/fclairamb/go-log/gokit"
)

func TestGoKit(_ *testing.T) {
	gkLogger := gklog.NewLogfmtLogger(gklog.NewSyncWriter(os.Stdout))
	logger := adapter.NewWrap(gkLogger)

	logger.Info("Hello world !")
}
