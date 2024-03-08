package samples

import (
	"testing"

	zerolog "github.com/rs/zerolog/log"

	adapter "github.com/fclairamb/go-log/zerolog"
)

func TestZeroLog(_ *testing.T) {
	logger := adapter.NewWrap(&zerolog.Logger)

	logger.Info("Hello world !")
}
