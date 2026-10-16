package hllogger

/*
	Journald prefixes are syslog values (see syslog.h)
	https://0pointer.de/blog/projects/journal-submit.html
*/

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
	case LevelDebug:
		return SystemdDebugPrefix
	case LevelInfo:
		return SystemdInfoPrefix
	case LevelNotice:
		return SystemdNoticePrefix
	case LevelWarning:
		return SystemdWarningPrefix
	case LevelError:
		return SystemdErrPrefix
	case LevelCritical:
		return SystemdCritPrefix
	case LevelAlert:
		return SystemdAlertPrefix
	case LevelEmergency:
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
