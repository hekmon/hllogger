package hllogger

import (
	"fmt"
	"os"
)

// Output is a wrapper to the log.Output() base fx
func (l *HlLogger) Output(message string) {
	l.logger.Output(2, message)
}

// Outputf is a wrapper to the the Output() fx but with formating support
func (l *HlLogger) Outputf(format string, a ...interface{}) {
	l.Output(fmt.Sprintf(format, a...))
}

// Debug is a wrapper to Output()
func (l *HlLogger) Debug(message string) {
	if l.llevel >= Debug {
		var prefix string
		if l.sysd {
			prefix = SystemdDebugPrefix
		}
		l.Output(prefix + "  DEBUG: " + message)
	}
}

// Debugf is a wrapper to Debug() adding formating support
func (l *HlLogger) Debugf(format string, a ...interface{}) {
	if l.llevel >= Debug {
		l.Debug(fmt.Sprintf(format, a...))
	}
}

// IsDebugShown returns true if the Debug level is in use
func (l *HlLogger) IsDebugShown() (shown bool) {
	return l.llevel >= Debug
}

// Info is a wrapper to Output()
func (l *HlLogger) Info(message string) {
	if l.llevel >= Info {
		var prefix string
		if l.sysd {
			prefix = SystemdInfoPrefix
		}
		l.Output(prefix + "   INFO: " + message)
	}
}

// Infof is a wrapper to Info() adding formating support
func (l *HlLogger) Infof(format string, a ...interface{}) {
	if l.llevel >= Info {
		l.Info(fmt.Sprintf(format, a...))
	}
}

// IsInfoShown returns true if the Info level is in use
func (l *HlLogger) IsInfoShown() (shown bool) {
	return l.llevel >= Info
}

// Warning is a wrapper to Output()
func (l *HlLogger) Warning(message string) {
	if l.llevel >= Warning {
		var prefix string
		if l.sysd {
			prefix = SystemdWarningPrefix
		}
		l.Output(prefix + "WARNING: " + message)
	}
}

// Warningf is a wrapper to Warning() adding formating support
func (l *HlLogger) Warningf(format string, a ...interface{}) {
	if l.llevel >= Warning {
		l.Warning(fmt.Sprintf(format, a...))
	}
}

// IsWarningShown returns true if the Warning level is in use
func (l *HlLogger) IsWarningShown() (shown bool) {
	return l.llevel >= Warning
}

// Error is a wrapper to Output()
func (l *HlLogger) Error(message string) {
	if l.llevel >= Error {
		var prefix string
		if l.sysd {
			prefix = SystemdErrPrefix
		}
		l.Output(prefix + "  ERROR: " + message)
	}
}

// Errorf is a wrapper to Error() adding formating support
func (l *HlLogger) Errorf(format string, a ...interface{}) {
	if l.llevel >= Error {
		l.Error(fmt.Sprintf(format, a...))
	}
}

// IsErrorShown returns true if the Error level is in use
func (l *HlLogger) IsErrorShown() (shown bool) {
	return l.llevel >= Error
}

// Fatal is a wrapper to Output()
// It will also print the message on stderr and exit with error code 1
func (l *HlLogger) Fatal(exitCode int, message string) {
	if l.llevel >= Fatal {
		var prefix string
		if l.sysd {
			prefix = SystemdCritPrefix
		}
		l.Output(prefix + "  FATAL: " + message)
	}
	os.Exit(exitCode)
}

// Fatalf is a wrapper to Fatal() adding formating support
func (l *HlLogger) Fatalf(exitCode int, format string, a ...interface{}) {
	if l.llevel >= Fatal {
		l.Fatal(exitCode, fmt.Sprintf(format, a...))
	}
	// in case l.llevel did not match, we still need to exit
	os.Exit(exitCode)
}

// IsFatalShown returns true if the Fatal level is in use
func (l *HlLogger) IsFatalShown() (shown bool) {
	return l.llevel >= Fatal
}
