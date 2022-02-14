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
	// Enable journald compat mode ?
	_, journaldPrefix = os.LookupEnv("INVOCATION_ID")
	if !(journaldPrefix && (output == os.Stdout || output == os.Stderr)) {
		// Switch to non systemd mode as conditions are not met
		journaldPrefix = false
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
