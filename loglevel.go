package hllogger

// LogLevel defines logging level constants
type LogLevel uint8

const (
	// Fatal level
	Fatal LogLevel = iota
	// Error level
	Error
	// Warning level
	Warning
	// Info level
	Info
	// Debug level
	Debug
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

const (
	// SystemdDebugPrefix is the string to prefix in Debug for systemd-journald
	SystemdDebugPrefix = "<7>"
	// SystemdInfoPrefix is the string to prefix in Info for systemd-journald
	SystemdInfoPrefix = "<6>"
	// SystemdNoticePrefix is the string to prefix in Notice for systemd-journald
	SystemdNoticePrefix = "<5>"
	// SystemdWarningPrefix is the string to prefix in Warning for systemd-journald
	SystemdWarningPrefix = "<4>"
	// SystemdErrPrefix is the string to prefix in Error for systemd-journald
	SystemdErrPrefix = "<3>"
	// SystemdCritPrefix is the string to prefix in Critical for systemd-journald
	SystemdCritPrefix = "<2>"
	// SystemdAlertPrefix is the string to prefix in Alert for systemd-journald
	SystemdAlertPrefix = "<1>"
	// SystemdEmergPrefix is the string to prefix in Emergency for systemd-journald
	SystemdEmergPrefix = "<0>"
)

// SystemdPrefix returns the appropriate prefix for systemd-journald
func (ll *LogLevel) SystemdPrefix() string {
	switch *ll {
	case Debug:
		return SystemdDebugPrefix
	case Info:
		return SystemdInfoPrefix
	case Warning:
		return SystemdWarningPrefix
	case Error:
		return SystemdErrPrefix
	case Fatal:
		return SystemdCritPrefix
	}
	return ""
}
