package rec

import (
	"os"
	"time"
)

var exitFn = os.Exit // nolint: gochecknoglobals

// Print outputs the log entry for the passed rec.Severity.
func (l *Logger) Print(severity Severity, message string, fields ...Field) {
	l.write(time.Now(), severity, message, fields...)
}

// Fatal outputs the log entry for the passed rec.Severity and call os.Exit(1).
func (l *Logger) Fatal(severity Severity, message string, fields ...Field) {
	l.write(time.Now(), severity, message, fields...)
	exitFn(1)
}

// Panic outputs the log entry for the passed rec.Severity and call panic(message).
func (l *Logger) Panic(severity Severity, message string, fields ...Field) {
	l.write(time.Now(), severity, message, fields...)
	panic(message)
}

// Default outputs the DEFAULT Severity log entry.
func (l *Logger) Default(message string, fields ...Field) {
	l.write(time.Now(), DEFAULT, message, fields...)
}

// Debug outputs the DEBUG Severity log entry.
func (l *Logger) Debug(message string, fields ...Field) {
	l.write(time.Now(), DEBUG, message, fields...)
}

// Info outputs the INFO Severity log entry.
func (l *Logger) Info(message string, fields ...Field) {
	l.write(time.Now(), INFO, message, fields...)
}

// Notice outputs the NOTICE Severity log entry.
func (l *Logger) Notice(message string, fields ...Field) {
	l.write(time.Now(), NOTICE, message, fields...)
}

// Warning outputs the WARNING Severity log entry.
func (l *Logger) Warning(message string, fields ...Field) {
	l.write(time.Now(), WARNING, message, fields...)
}

// Error outputs the ERROR Severity log entry.
func (l *Logger) Error(message string, fields ...Field) {
	l.write(time.Now(), ERROR, message, fields...)
}

// Critical outputs the CRITICAL Severity log entry.
func (l *Logger) Critical(message string, fields ...Field) {
	l.write(time.Now(), CRITICAL, message, fields...)
}

// Alert outputs the ALERT Severity log entry.
func (l *Logger) Alert(message string, fields ...Field) {
	l.write(time.Now(), ALERT, message, fields...)
}

// Emergency outputs the EMERGENCY Severity log entry.
func (l *Logger) Emergency(message string, fields ...Field) {
	l.write(time.Now(), EMERGENCY, message, fields...)
}
