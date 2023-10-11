package hllogger

import "fmt"

// LogLevel defines logging level constants
type LogLevel uint8

const (
	// Emergency level
	LevelEmergency LogLevel = iota
	// Alert level
	LevelAlert
	// Critical level
	LevelCritical
	// Error level
	LevelError
	// Warning level
	LevelWarning
	// Notice level
	LevelNotice
	// Info level
	LevelInfo
	// Debug level
	LevelDebug
)

func (ll LogLevel) String() string {
	switch ll {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelNotice:
		return "NOTICE"
	case LevelWarning:
		return "WARNING"
	case LevelError:
		return "ERROR"
	case LevelCritical:
		return "CRITICAL"
	case LevelAlert:
		return "ALERT"
	case LevelEmergency:
		return "EMERGENCY"
	}
	return "<unknown>"
}

// GoString implements the GoStringer interface:
// https://pkg.go.dev/fmt#GoStringer
func (ll LogLevel) GoString() string {
	return fmt.Sprintf("%s(%d)", ll, ll)
}
