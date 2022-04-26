package samples

import (
	"testing"

	"github.com/inconshreveable/log15"

	adapter "github.com/fclairamb/go-log/log15"
)

func TestLog15(_ *testing.T) {
	log15Logger := log15.New()
	logger := adapter.NewWrap(log15Logger)

	logger.Info("Hello world !")
}
