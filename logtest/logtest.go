// Package logtest provides testing code for loggers
package logtest

// https://stackoverflow.com/a/31794241

import (
	"testing"

	"github.com/fclairamb/go-log"
)

// TestLogger is shared test function
func TestLogger(t *testing.T, logger log.Logger) {
	logger = logger.With("test", t.Name())
	logger.Debug("This is debug", "key", "val")
	logger.Info("This is info", "key", "val")
	logger.Warn("This is warning", "key", "val")
	logger.Error("This is error", "key", "val")

	// Buggy logging
	logger.Info("Event", "key1", "val1", "key2")
}
