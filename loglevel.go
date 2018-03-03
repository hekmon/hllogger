package hllogger

// LogLevel defines logging level constants
type LogLevel uint8

const (
	// Debug level
	Debug LogLevel = iota
	// Info level
	Info
	// Warning level
	Warning
	// Error level
	Error
	// Fatal level
	Fatal
)

func (ll *LogLevel) String() string {
	switch *ll {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	}
	return "<unknown>"
}
