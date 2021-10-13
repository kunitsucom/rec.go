package rec

import (
	"fmt"
	"time"
)

type formatLogger struct {
	l *Logger
}

// F returns format logger utilities.
func (l *Logger) F() *formatLogger { // nolint: revive
	return &formatLogger{
		l: l,
	}
}

// F returns default `*rec.Logger` in rec package as `*rec.formatLogger`.
func F() *formatLogger { // nolint: revive
	return defaultLogger.F()
}

// Print outputs the log entry for the passed rec.Severity.
func (f *formatLogger) Printf(severity Severity, format string, v ...interface{}) {
	f.l.write(time.Now(), severity, fmt.Sprintf(format, v...))
}

// Fatal outputs the log entry for the passed rec.Severity and call os.Exit(1).
func (f *formatLogger) Fatalf(severity Severity, format string, v ...interface{}) {
	f.l.write(time.Now(), severity, fmt.Sprintf(format, v...))
	exitFn(1)
}

// Panic outputs the log entry for the passed rec.Severity and call panic(message).
func (f *formatLogger) Panicf(severity Severity, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	f.l.write(time.Now(), severity, message)
	panic(message)
}

// Default outputs the DEFAULT Severity log entry.
func (f *formatLogger) Defaultf(format string, v ...interface{}) {
	f.l.write(time.Now(), DEFAULT, fmt.Sprintf(format, v...))
}

// Debug outputs the DEBUG Severity log entry.
func (f *formatLogger) Debugf(format string, v ...interface{}) {
	f.l.write(time.Now(), DEBUG, fmt.Sprintf(format, v...))
}

// Info outputs the INFO Severity log entry.
func (f *formatLogger) Infof(format string, v ...interface{}) {
	f.l.write(time.Now(), INFO, fmt.Sprintf(format, v...))
}

// Notice outputs the NOTICE Severity log entry.
func (f *formatLogger) Noticef(format string, v ...interface{}) {
	f.l.write(time.Now(), NOTICE, fmt.Sprintf(format, v...))
}

// Warning outputs the WARNING Severity log entry.
func (f *formatLogger) Warningf(format string, v ...interface{}) {
	f.l.write(time.Now(), WARNING, fmt.Sprintf(format, v...))
}

// Error outputs the ERROR Severity log entry.
func (f *formatLogger) Errorf(format string, v ...interface{}) {
	f.l.write(time.Now(), ERROR, fmt.Sprintf(format, v...))
}

// Critical outputs the CRITICAL Severity log entry.
func (f *formatLogger) Criticalf(format string, v ...interface{}) {
	f.l.write(time.Now(), CRITICAL, fmt.Sprintf(format, v...))
}

// Alert outputs the ALERT Severity log entry.
func (f *formatLogger) Alertf(format string, v ...interface{}) {
	f.l.write(time.Now(), ALERT, fmt.Sprintf(format, v...))
}

// Emergency outputs the EMERGENCY Severity log entry.
func (f *formatLogger) Emergencyf(format string, v ...interface{}) {
	f.l.write(time.Now(), EMERGENCY, fmt.Sprintf(format, v...))
}
