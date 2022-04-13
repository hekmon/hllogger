package hllogger

import (
	"fmt"
	"io"
	"log"
	"os"
)

// New is the constructor for a Logger object
func New(output io.Writer, level LogLevel) *Logger {
	// Detect special compatibility modes
	_, systemdInvocation := os.LookupEnv("INVOCATION_ID")
	lambdaFunctionName, awsLambdaFunction := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME")
	// Prepare configuration accordingly
	if systemdInvocation && output != os.Stdout && output != os.Stderr {
		// launched by systemd but logger is not being redirected to std output (may be to a file ?)
		// disabling journald compat mode
		systemdInvocation = false
	}
	var flags int
	if !systemdInvocation && !awsLambdaFunction {
		flags = log.Ltime | log.Ldate
	}
	var prefix string
	if awsLambdaFunction {
		prefix = fmt.Sprintf("[%s] ", lambdaFunctionName)
	}
	// Return the initialized logger
	return &Logger{
		journald: systemdInvocation,
		llevel:   level,
		logger:   log.New(output, prefix, flags),
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

// SetOutput sets the output destination for this logger instance.
func (l *Logger) SetOutput(w io.Writer) {
	l.logger.SetOutput(w)
}

// SetLogLevel sets the log level for this logger instance.
// This function is not thread safe, do not set log level if you are printing logs already.
func (l *Logger) SetLogLevel(ll LogLevel) {
	l.llevel = ll
}

// SetFlags sets the logger flags for this logger instance.
func (l *Logger) SetFlags(flags int) {
	l.logger.SetFlags(flags)
}
