package level

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevels(t *testing.T) {
	a := assert.New(t)

	type levelDef struct {
		name string
		lvl  Level
	}

	levels := []levelDef{
		{"debug", Debug},
		{"info", Info},
		{"warn", Warning},
		{"error", Error},
		{"panic", Panic},
	}

	for _, l := range levels {
		a.Equal(l.name, l.lvl.String())
		a.Equal(l.lvl, FromString(l.name))
	}
}

func TestShouldLog(t *testing.T) {
	a := assert.New(t)

	a.True(Debug.ShouldLog(Debug))
	a.True(Debug.ShouldLog(Info))
	a.True(Debug.ShouldLog(Warning))
	a.True(Debug.ShouldLog(Error))
	a.True(Debug.ShouldLog(Panic))
	a.False(Warning.ShouldLog(Debug))
	a.False(Warning.ShouldLog(Info))
	a.True(Warning.ShouldLog(Warning))
	a.True(Warning.ShouldLog(Error))
	a.True(Warning.ShouldLog(Panic))
}
