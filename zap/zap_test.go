package log15

import (
	"errors"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/fclairamb/go-log/logtest"
)

func TestZap(t *testing.T) {
	a := require.New(t)

	logger, err := New()
	a.NoError(err)

	logtest.TestLogger(t, logger)
}

var ErrLogger = errors.New("logger error")

func TestFailingZap(t *testing.T) {
	a := require.New(t)

	logger, err := New(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		if err := core.Write(zapcore.Entry{}, nil); err != nil { // nolint: exhaustivestruct
			log.Fatal(err)
		}

		return core
	}))
	a.Error(err)
	a.Nil(logger)

	logtest.TestLogger(t, logger)
}
