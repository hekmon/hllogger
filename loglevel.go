package hllogger

import "fmt"

// LogLevel defines logging level constants
type LogLevel uint8

const (
	// Emergency level
	Emergency LogLevel = iota
	// Alert level
	Alert
	// Critical level
	Critical
	// Error level
	Error
	// Warning level
	Warning
	// Notice level
	Notice
	// Info level
	Info
	// Debug level
	Debug
)

func (ll LogLevel) String() string {
	switch ll {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Notice:
		return "NOTICE"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	case Critical:
		return "CRITICAL"
	case Alert:
		return "ALERT"
	case Emergency:
		return "EMERGENCY"
	}
	return "<unknown>"
}

// GoString implements the GoStringer interface:
// https://pkg.go.dev/fmt#GoStringer
func (ll LogLevel) GoString() string {
	return fmt.Sprintf("%s(%d)", ll, ll)
}
