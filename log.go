package hllogger

import (
	"fmt"
)

/*
	Output
*/

// Output is a wrapper to the log.Output() base fx
func (l *Logger) Output(message string) {
	l.logger.Output(2, message)
}

// Outputf is a wrapper to the the Output() fx but with formating support
func (l *Logger) Outputf(format string, a ...interface{}) {
	l.Output(fmt.Sprintf(format, a...))
}

/*
	Debug
*/

// Debug logs a message at the Debug level
func (l *Logger) Debug(message string) {
	if l.llevel >= LevelDebug {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(LevelDebug), LevelDebug, message)
	}
}

// Debugf logs a message at the Debug level with formating support
func (l *Logger) Debugf(format string, a ...interface{}) {
	if l.llevel >= LevelDebug {
		l.Debug(fmt.Sprintf(format, a...))
	}
}

// IsDebugShown returns true if the Debug level is in use
func (l *Logger) IsDebugShown() (shown bool) {
	return l.llevel >= LevelDebug
}

/*
	Info
*/

// Info logs a message at the Info level
func (l *Logger) Info(message string) {
	if l.llevel >= LevelInfo {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(LevelInfo), LevelInfo, message)
	}
}

// Infof logs a message at the Info level with formating support
func (l *Logger) Infof(format string, a ...interface{}) {
	if l.llevel >= LevelInfo {
		l.Info(fmt.Sprintf(format, a...))
	}
}

// IsInfoShown returns true if the Info level is in use
func (l *Logger) IsInfoShown() (shown bool) {
	return l.llevel >= LevelInfo
}

/*
	Notice
*/

// Notice logs a message at the Notice level
func (l *Logger) Notice(message string) {
	if l.llevel >= LevelNotice {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(LevelNotice), LevelNotice, message)
	}
}

// Noticef logs a message at the Notice level with formating support
func (l *Logger) Noticef(format string, a ...interface{}) {
	if l.llevel >= LevelNotice {
		l.Notice(fmt.Sprintf(format, a...))
	}
}

// IsNoticeShown returns true if the Notice level is in use
func (l *Logger) IsNoticeShown() (shown bool) {
	return l.llevel >= LevelNotice
}

/*
   Warning
*/

// Warning logs a message at the Warning level
func (l *Logger) Warning(message string) {
	if l.llevel >= LevelWarning {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(LevelWarning), LevelWarning, message)
	}
}

// Warningf logs a message at the Warning level with formating support
func (l *Logger) Warningf(format string, a ...interface{}) {
	if l.llevel >= LevelWarning {
		l.Warning(fmt.Sprintf(format, a...))
	}
}

// IsWarningShown returns true if the Warning level is in use
func (l *Logger) IsWarningShown() (shown bool) {
	return l.llevel >= LevelWarning
}

/*
   Error
*/

// Error logs a message at the Error level
func (l *Logger) Error(message string) {
	if l.llevel >= LevelError {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(LevelError), LevelError, message)
	}
}

// Errorf logs a message at the Error level with formating support
func (l *Logger) Errorf(format string, a ...interface{}) {
	if l.llevel >= LevelError {
		l.Error(fmt.Sprintf(format, a...))
	}
}

// IsErrorShown returns true if the Error level is in use
func (l *Logger) IsErrorShown() (shown bool) {
	return l.llevel >= LevelError
}

/*
   Critical
*/

// Critical logs a message at the Critical level
func (l *Logger) Critical(message string) {
	if l.llevel >= LevelCritical {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(LevelCritical), LevelCritical, message)
	}
}

// Criticalf logs a message at the Critical level with formating support
func (l *Logger) Criticalf(format string, a ...interface{}) {
	if l.llevel >= LevelCritical {
		l.Critical(fmt.Sprintf(format, a...))
	}
}

// IsCriticalShown returns true if the Critical level is in use
func (l *Logger) IsCriticalShown() (shown bool) {
	return l.llevel >= LevelCritical
}

/*
   Alert
*/

// Alert logs a message at the Alert level
func (l *Logger) Alert(message string) {
	if l.llevel >= LevelAlert {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(LevelAlert), LevelAlert, message)
	}
}

// Alertf logs a message at the Alert level with formating support
func (l *Logger) Alertf(format string, a ...interface{}) {
	if l.llevel >= LevelAlert {
		l.Alert(fmt.Sprintf(format, a...))
	}
}

// IsAlertShown returns true if the Alert level is in use
func (l *Logger) IsAlertShown() (shown bool) {
	return l.llevel >= LevelAlert
}

/*
   Emergency
*/

// Emergency logs a message at the XXX level
func (l *Logger) Emergency(message string) {
	if l.llevel >= LevelEmergency {
		l.Outputf("%s%9s: %s", l.getJournaldPrefix(LevelEmergency), LevelEmergency, message)
	}
}

// Emergencyf logs a message at the Emergency level with formating support
func (l *Logger) Emergencyf(format string, a ...interface{}) {
	if l.llevel >= LevelEmergency {
		l.Emergency(fmt.Sprintf(format, a...))
	}
}

// IsEmergencyShown returns true if the Emergency level is in use
func (l *Logger) IsEmergencyShown() (shown bool) {
	return l.llevel >= LevelEmergency
}
