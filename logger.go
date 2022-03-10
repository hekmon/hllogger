package hllogger

import (
	"io"
	"log"
	"os"
)

// New is the constructor for a Logger object
func New(output io.Writer, level LogLevel) *Logger {
	var (
		systemdInvocation bool
		awsLogGroup       bool
		flags             int
	)
	// Detect special compatibility modes
	_, systemdInvocation = os.LookupEnv("INVOCATION_ID")
	_, awsLogGroup = os.LookupEnv("AWS_LAMBDA_LOG_GROUP_NAME")
	// Prepare configuration accordingly
	if systemdInvocation && output != os.Stdout && output != os.Stderr {
		// launched by systemd but logger is not being redirected to std output (may be to a file ?)
		// disabling journald compat mode
		systemdInvocation = false
	}
	if !systemdInvocation && !awsLogGroup {
		flags = log.Ltime | log.Ldate
	}
	// Return the initialized logger
	return &Logger{
		journald: systemdInvocation,
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
