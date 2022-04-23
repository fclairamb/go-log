package samples

import (
	"testing"

	adapter "github.com/fclairamb/go-log/zerolog"

	zerolog "github.com/rs/zerolog/log"
)

func TestZeroLog(t *testing.T) {
	logger := adapter.NewWrap(&zerolog.Logger)

	logger.Info("Hello world !")
}
