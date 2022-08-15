// Package logtest provides testing code for loggers
package logtest

// https://stackoverflow.com/a/31794241

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fclairamb/go-log"
)

// TestLogger is shared test function
func TestLogger(t *testing.T, logger log.Logger) {
	t.Helper()

	logger = logger.With("test", t.Name())
	t.Run("levels", func(t *testing.T) {
		logger.Debug("This is debug", "key", "val")
		logger.Info("This is info", "key", "val")
		logger.Warn("This is warning", "key", "val")
		logger.Error("This is error", "key", "val")

		assert.Panics(t, func() {
			logger.Panic("This is panic", "key1", "val1", "key2", "val2")
		})
	})

	// Buggy logging
	t.Run("buggy", func(t *testing.T) {
		logger.Info("Buggy log", "key1", "val1", "key2")
	})

	t.Run("with", func(t *testing.T) {
		subLogger := logger.With("key", "val")
		subLogger.Info("This is info")
	})

	t.Run("types", func(t *testing.T) {
		logger.Info(
			"Types",
			"string", "string",
			"int", 1,
			"int8", int8(1),
			"int16", int16(1),
			"int32", int32(1),
			"int64", int64(1),
			"uint", uint(1),
			"uint8", uint8(1),
			"uint16", uint16(1),
			"uint32", uint32(1),
			"uint64", uint64(1),
			"float32", float32(1),
			"float64", float64(1),
			"bool", true,
			"byte", byte(1),
			"rune", rune(1),
			"interface", interface{}(1),
			"error1", nil,
			"error2", os.ErrInvalid,
			// "complex64", complex64(1),
			// "complex128", complex128(1),
		)
	})
}
