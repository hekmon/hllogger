package hllogger

import (
	"fmt"
)

/*
	Output
*/

// Output is a wrapper to the log.Output() base fx
func (l *HlLogger) Output(message string) {
	l.logger.Output(2, message)
}

// Outputf is a wrapper to the the Output() fx but with formating support
func (l *HlLogger) Outputf(format string, a ...interface{}) {
	l.Output(fmt.Sprintf(format, a...))
}

/*
	Debug
*/

// Debug is a wrapper to Output()
func (l *HlLogger) Debug(message string) {
	if l.llevel >= Debug {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(Debug), Debug, message)
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

/*
	Info
*/

// Info is a wrapper to Output()
func (l *HlLogger) Info(message string) {
	if l.llevel >= Info {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(Info), Info, message)
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

/*
	Notice
*/

// Notice is a wrapper to Output()
func (l *HlLogger) Notice(message string) {
	if l.llevel >= Notice {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(Notice), Notice, message)
	}
}

// Noticef is a wrapper to Notice() adding formating support
func (l *HlLogger) Noticef(format string, a ...interface{}) {
	if l.llevel >= Notice {
		l.Notice(fmt.Sprintf(format, a...))
	}
}

// IsNoticeShown returns true if the Notice level is in use
func (l *HlLogger) IsNoticeShown() (shown bool) {
	return l.llevel >= Notice
}

/*
   Warning
*/

// Warning is a wrapper to Output()
func (l *HlLogger) Warning(message string) {
	if l.llevel >= Warning {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(Warning), Warning, message)
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

/*
   Error
*/

// Error is a wrapper to Output()
func (l *HlLogger) Error(message string) {
	if l.llevel >= Error {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(Error), Error, message)
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

/*
   Critical
*/

// Critical is a wrapper to Output()
func (l *HlLogger) Critical(message string) {
	if l.llevel >= Critical {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(Critical), Critical, message)
	}
}

// Criticalf is a wrapper to Critical() adding formating support
func (l *HlLogger) Criticalf(format string, a ...interface{}) {
	if l.llevel >= Critical {
		l.Critical(fmt.Sprintf(format, a...))
	}
}

// IsCriticalShown returns true if the Critical level is in use
func (l *HlLogger) IsCriticalShown() (shown bool) {
	return l.llevel >= Critical
}

/*
   Alert
*/

// Alert is a wrapper to Output()
func (l *HlLogger) Alert(message string) {
	if l.llevel >= Alert {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(Alert), Alert, message)
	}
}

// Alertf is a wrapper to Alert() adding formating support
func (l *HlLogger) Alertf(format string, a ...interface{}) {
	if l.llevel >= Alert {
		l.Alert(fmt.Sprintf(format, a...))
	}
}

// IsAlertShown returns true if the Alert level is in use
func (l *HlLogger) IsAlertShown() (shown bool) {
	return l.llevel >= Alert
}

/*
   Emergency
*/

// Emergency is a wrapper to Output()
func (l *HlLogger) Emergency(message string) {
	if l.llevel >= Emergency {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(Emergency), Emergency, message)
	}
}

// Emergencyf is a wrapper to Emergency() adding formating support
func (l *HlLogger) Emergencyf(format string, a ...interface{}) {
	if l.llevel >= Emergency {
		l.Emergency(fmt.Sprintf(format, a...))
	}
}

// IsEmergencyShown returns true if the Emergency level is in use
func (l *HlLogger) IsEmergencyShown() (shown bool) {
	return l.llevel >= Emergency
}
