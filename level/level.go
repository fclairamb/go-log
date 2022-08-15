// Package level allows to do level filtering on go-log side
package level

// Level defines a log level
type Level uint8

const (
	Debug Level = iota // Debug level
	Info               // Info level
	Warn               // Warning level
	Error              // Error level
	Panic              // Panic level
)

func (l Level) String() string {
	switch l {
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case Panic:
		return "panic"
	default:
		return "unknown"
	}
}

// FromString returns the level from a string
func FromString(s string) Level {
	switch s {
	case "debug":
		return Debug
	case "info":
		return Info
	case "warn":
		return Warn
	case "error":
		return Error
	case "panic":
		return Panic
	default:
		panic("unknown level")
	}
}

// ShouldLog returns true if the level is greater than the given level
func (l Level) ShouldLog(level Level) bool {
	return l <= level
}
