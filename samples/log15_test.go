package samples

import (
	"testing"

	"github.com/inconshreveable/log15"

	adapter "github.com/fclairamb/go-log/log15"
	"github.com/fclairamb/go-log/logtest"
)

func TestLog15(t *testing.T) {
	log15Logger := log15.New()
	logger := adapter.NewWrap(log15Logger)

	logtest.TestLogger(t, logger)
}
