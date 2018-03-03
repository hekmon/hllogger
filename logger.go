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
	llevel LogLevel
	logger *log.Logger
}

// New is the constructor for a HlLogger object
func New(output io.Writer, prefix string, logLevel LogLevel, loggerFlags int) *HlLogger {
	// Basic init
	return &HlLogger{
		llevel: logLevel,
		logger: log.New(output, prefix, loggerFlags),
	}
}
