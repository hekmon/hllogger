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
	journald bool
	llevel   LogLevel
	logger   *log.Logger
}

// Config allows to configure an hllogger
type Config struct {
	LogLevel       LogLevel
	LoggerFlags    int
	JournaldPrefix bool
}

// New is the constructor for a HlLogger object
func New(output io.Writer, c Config) *HlLogger {
	return &HlLogger{
		journald: c.JournaldPrefix,
		llevel:   c.LogLevel,
		logger:   log.New(output, "", c.LoggerFlags),
	}
}
