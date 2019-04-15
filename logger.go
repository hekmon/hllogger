package hllogger

import (
	"io"
	"log"
)

// Forwarders to the log package flags
const (
	Ldate         = log.Ldate
	Ltime         = log.Ltime
	Lmicroseconds = log.Lmicroseconds
	Llongfile     = log.Llongfile
	Lshortfile    = log.Lshortfile
	LUTC          = log.LUTC
	LstdFlags     = log.LstdFlags
)

// HlLogger is a wrapper around the log/logger
type HlLogger struct {
	sysd   bool
	llevel LogLevel
	logger *log.Logger
}

// Config allows to configure an hllogger
type Config struct {
	LogLevel              LogLevel
	LoggerFlags           int
	SystemdJournaldCompat bool
}

// New is the constructor for a HlLogger object
func New(output io.Writer, c *Config) *HlLogger {
	if c == nil {
		c = &Config{
			LogLevel: Info,
		}
	}
	return &HlLogger{
		sysd:   c.SystemdJournaldCompat,
		llevel: c.LogLevel,
		logger: log.New(output, "", c.LoggerFlags),
	}
}
