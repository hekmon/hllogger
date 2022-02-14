package hllogger

import (
	"io"
	"log"
	"os"
)

// New is the constructor for a Logger object
func New(output io.Writer, level LogLevel) *Logger {
	var (
		journaldPrefix bool
		flags          int
	)
	// Is systemd detected ?
	_, journaldPrefix = os.LookupEnv("INVOCATION_ID")
	if !journaldPrefix {
		// systemd not detected, add date and time as prefix for logging
		flags = log.Ltime | log.Ldate
	}
	// Return the initialized logger
	return &Logger{
		journald: journaldPrefix,
		llevel:   level,
		logger:   log.New(output, "", flags),
	}
}

// Logger is a wrapper around the log/logger
type Logger struct {
	journald bool
	llevel   LogLevel
	logger   *log.Logger
}

// GetStdLogger returns the lower level std logger used by this instance of Logger
func (l *Logger) GetStdLogger() *log.Logger {
	return l.logger
}
