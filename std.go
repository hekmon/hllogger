package hllogger

import (
	"io"
	"os"
)

// std logger is a *Logger that outputs to stdout at info level by default.
// You can configure it with Set* methods.
var std = New(os.Stdout, LevelInfo)

// Default returns the default standard logger used by the package-level output functions.
func Default() *Logger {
	return std
}

// SetOutput sets the output destination for the default standard logger.
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

// SetLogLevel sets the log level for the default standard logger.
// This function is not thread safe, do not set log level if you are printing logs already.
func SetLogLevel(l LogLevel) {
	std.SetLogLevel(l)
}

// SetFlags sets the logger flags for the default standard logger.
func SetFlags(flags int) {
	std.SetFlags(flags)
}

// Debug logs a message at the Debug level with the default standard logger.
func Debug(message string) {
	std.Debug(message)
}

// Debugf logs a formatted message at the Debug level with the default standard logger.
func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}

// IsDebugShown returns true if the Debug level is in use on the default standard logger.
func IsDebugShown() bool {
	return std.IsDebugShown()
}

// Info logs a message at the Info level with the default standard logger.
func Info(message string) {
	std.Info(message)
}

// Infof logs a formatted message at the Info level with the default standard logger.
func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}

// IsInfoShown returns true if the Info level is in use on the default standard logger.
func IsInfoShown() bool {
	return std.IsInfoShown()
}

// Notice logs a message at the Notice level with the default standard logger.
func Notice(message string) {
	std.Notice(message)
}

// Noticef logs a formatted message at the Notice level with the default standard logger.
func Noticef(format string, v ...interface{}) {
	std.Noticef(format, v...)
}

// IsNoticeShown returns true if the Notice level is in use on the default standard logger.
func IsNoticeShown() bool {
	return std.IsNoticeShown()
}

// Warning logs a message at the Warning level with the default standard logger.
func Warning(message string) {
	std.Warning(message)
}

// Warningf logs a formatted message at the Warning level with the default standard logger.
func Warningf(format string, v ...interface{}) {
	std.Warningf(format, v...)
}

// IsWarningShown returns true if the Warning level is in use on the default standard logger.
func IsWarningShown() bool {
	return std.IsWarningShown()
}

// Error logs a message at the Error level with the default standard logger.
func Error(message string) {
	std.Error(message)
}

// Errorf logs a formatted message at the Error level with the default standard logger.
func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

// IsErrorShown returns true if the Error level is in use on the default standard logger.
func IsErrorShown() bool {
	return std.IsErrorShown()
}

// Critical logs a message at the Critical level with the default standard logger.
func Critical(message string) {
	std.Critical(message)
}

// Criticalf logs a formatted message at the Critical level with the default standard logger.
func Criticalf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

// IsCriticalShown returns true if the Critical level is in use on the default standard logger.
func IsCriticalShown() bool {
	return std.IsCriticalShown()
}

// Alert logs a message at the Alert level with the default standard logger.
func Alert(message string) {
	std.Alert(message)
}

// Alertf logs a formatted message at the Alert level with the default standard logger.
func Alertf(format string, v ...interface{}) {
	std.Alertf(format, v...)
}

// IsAlertShown returns true if the Alert level is in use on the default standard logger.
func IsAlertShown() bool {
	return std.IsAlertShown()
}

// Emergency logs a message at the Emergency level with the default standard logger.
func Emergency(message string) {
	std.Emergency(message)
}

// Emergencyf logs a formatted message at the Emergency level with the default standard logger.
func Emergencyf(format string, v ...interface{}) {
	std.Emergencyf(format, v...)
}

// IsEmergencyShown returns true if the Emergency level is in use on the default standard logger.
func IsEmergencyShown() bool {
	return std.IsEmergencyShown()
}
