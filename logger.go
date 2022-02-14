package hllogger

import (
	"io"
	"log"
	"os"
)

// New is the constructor for a HlLogger object
func New(output io.Writer, level LogLevel) *HlLogger {
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
	return &HlLogger{
		journald: journaldPrefix,
		llevel:   level,
		logger:   log.New(output, "", flags),
	}
}

// HlLogger is a wrapper around the log/logger
type HlLogger struct {
	journald bool
	llevel   LogLevel
	logger   *log.Logger
}
