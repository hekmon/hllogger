package hllogger

const (
	// SystemdEmergPrefix is the string to prefix in Emergency for systemd-journald
	SystemdEmergPrefix = "<0>"
	// SystemdAlertPrefix is the string to prefix in Alert for systemd-journald
	SystemdAlertPrefix = "<1>"
	// SystemdCritPrefix is the string to prefix in Critical for systemd-journald
	SystemdCritPrefix = "<2>"
	// SystemdErrPrefix is the string to prefix in Error for systemd-journald
	SystemdErrPrefix = "<3>"
	// SystemdWarningPrefix is the string to prefix in Warning for systemd-journald
	SystemdWarningPrefix = "<4>"
	// SystemdNoticePrefix is the string to prefix in Notice for systemd-journald
	SystemdNoticePrefix = "<5>"
	// SystemdInfoPrefix is the string to prefix in Info for systemd-journald
	SystemdInfoPrefix = "<6>"
	// SystemdDebugPrefix is the string to prefix in Debug for systemd-journald
	SystemdDebugPrefix = "<7>"
)

// SystemdPrefix returns the appropriate prefix for systemd-journald
func (ll LogLevel) SystemdPrefix() string {
	switch ll {
	case Debug:
		return SystemdDebugPrefix
	case Info:
		return SystemdInfoPrefix
	case Notice:
		return SystemdNoticePrefix
	case Warning:
		return SystemdWarningPrefix
	case Error:
		return SystemdErrPrefix
	case Critical:
		return SystemdCritPrefix
	case Alert:
		return SystemdAlertPrefix
	case Emergency:
		return SystemdEmergPrefix
	default:
		return ""
	}
}

func (hll *Logger) getJournaldPrefix(ll LogLevel) string {
	if hll.journald {
		return ll.SystemdPrefix()
	}
	return ""
}
