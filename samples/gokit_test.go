package samples

import (
	"os"
	"testing"

	gklog "github.com/go-kit/log"

	adapter "github.com/fclairamb/go-log/gokit"
	"github.com/fclairamb/go-log/logtest"
)

func TestGoKit(t *testing.T) {
	gkLogger := gklog.NewLogfmtLogger(gklog.NewSyncWriter(os.Stdout))
	logger := adapter.NewWrap(gkLogger)

	logtest.TestLogger(t, logger)
}
