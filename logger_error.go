package rec

import (
	"errors"
	"time"
)

type errorLogger struct {
	l *Logger
}

// E returns error logger utilities.
func (l *Logger) E() *errorLogger { // nolint: revive
	return &errorLogger{
		l: l,
	}
}

// E returns default `*rec.Logger` in rec package as `*rec.errorLogger`.
func E() *errorLogger { // nolint: revive
	return defaultLogger.E()
}

type errorReturner struct {
	e error
}

// Err returns error.
func (er *errorReturner) Err() error {
	return er.e
}

const formattedNilString = "<nil>"

// Print is alias of below:
//
//	Print(severity, err.Error(), rec.Error(err), rec.ErrorStacktrace(err))
func (e *errorLogger) Print(severity Severity, err error, fields ...Field) *errorReturner {
	var message string

	if errors.Is(err, nil) {
		message = formattedNilString
	} else {
		message = err.Error()
	}

	e.l.write(time.Now(), severity, message, append([]Field{Error(err), ErrorStacktrace(err)}, fields...)...)

	return &errorReturner{err}
}

// Fatal is alias of below:
//
//	Fatal(severity, err.Error(), rec.Error(err), rec.ErrorStacktrace(err))
func (e *errorLogger) Fatal(severity Severity, err error, fields ...Field) {
	var message string

	if errors.Is(err, nil) {
		message = formattedNilString
	} else {
		message = err.Error()
	}

	e.l.write(time.Now(), severity, message, append([]Field{Error(err), ErrorStacktrace(err)}, fields...)...)

	exitFn(1)
}

// Panic is alias of below:
//
//	Panic(severity, err.Error(), rec.Error(err), rec.ErrorStacktrace(err))
func (e *errorLogger) Panic(severity Severity, err error, fields ...Field) {
	var message string

	if errors.Is(err, nil) {
		message = formattedNilString
	} else {
		message = err.Error()
	}

	e.l.write(time.Now(), severity, message, append([]Field{Error(err), ErrorStacktrace(err)}, fields...)...)

	panic(err)
}

// Default is alias of below:
//
//	Default(err.Error(), rec.Error(err), rec.ErrorStacktrace(err))
func (e *errorLogger) Default(err error, fields ...Field) *errorReturner {
	var message string

	if errors.Is(err, nil) {
		message = formattedNilString
	} else {
		message = err.Error()
	}

	e.l.write(time.Now(), DEFAULT, message, append([]Field{Error(err), ErrorStacktrace(err)}, fields...)...)

	return &errorReturner{err}
}

// Debug is alias of below:
//
//	Debug(err.Error(), rec.Error(err), rec.ErrorStacktrace(err))
func (e *errorLogger) Debug(err error, fields ...Field) *errorReturner {
	var message string

	if errors.Is(err, nil) {
		message = formattedNilString
	} else {
		message = err.Error()
	}

	e.l.write(time.Now(), DEBUG, message, append([]Field{Error(err), ErrorStacktrace(err)}, fields...)...)

	return &errorReturner{err}
}

// Info is alias of below:
//
//	Info(err.Error(), rec.Error(err), rec.ErrorStacktrace(err))
func (e *errorLogger) Info(err error, fields ...Field) *errorReturner {
	var message string

	if errors.Is(err, nil) {
		message = formattedNilString
	} else {
		message = err.Error()
	}

	e.l.write(time.Now(), INFO, message, append([]Field{Error(err), ErrorStacktrace(err)}, fields...)...)

	return &errorReturner{err}
}

// Notice is alias of below:
//
//	Notice(err.Error(), rec.Error(err), rec.ErrorStacktrace(err))
func (e *errorLogger) Notice(err error, fields ...Field) *errorReturner {
	var message string

	if errors.Is(err, nil) {
		message = formattedNilString
	} else {
		message = err.Error()
	}

	e.l.write(time.Now(), NOTICE, message, append([]Field{Error(err), ErrorStacktrace(err)}, fields...)...)

	return &errorReturner{err}
}

// Warning is alias of below:
//
//	Warning(err.Error(), rec.Error(err), rec.ErrorStacktrace(err))
func (e *errorLogger) Warning(err error, fields ...Field) *errorReturner {
	var message string

	if errors.Is(err, nil) {
		message = formattedNilString
	} else {
		message = err.Error()
	}

	e.l.write(time.Now(), WARNING, message, append([]Field{Error(err), ErrorStacktrace(err)}, fields...)...)

	return &errorReturner{err}
}

// Error is alias of below:
//
//	Error(err.Error(), rec.Error(err), rec.ErrorStacktrace(err))
func (e *errorLogger) Error(err error, fields ...Field) *errorReturner {
	var message string

	if errors.Is(err, nil) {
		message = formattedNilString
	} else {
		message = err.Error()
	}

	e.l.write(time.Now(), ERROR, message, append([]Field{Error(err), ErrorStacktrace(err)}, fields...)...)

	return &errorReturner{err}
}

// Critical is alias of below:
//
//	Critical(err.Error(), rec.Error(err), rec.ErrorStacktrace(err))
func (e *errorLogger) Critical(err error, fields ...Field) *errorReturner {
	var message string

	if errors.Is(err, nil) {
		message = formattedNilString
	} else {
		message = err.Error()
	}

	e.l.write(time.Now(), CRITICAL, message, append([]Field{Error(err), ErrorStacktrace(err)}, fields...)...)

	return &errorReturner{err}
}

// Alert is alias of below:
//
//	Alert(err.Error(), rec.Error(err), rec.ErrorStacktrace(err))
func (e *errorLogger) Alert(err error, fields ...Field) *errorReturner {
	var message string

	if errors.Is(err, nil) {
		message = formattedNilString
	} else {
		message = err.Error()
	}

	e.l.write(time.Now(), ALERT, message, append([]Field{Error(err), ErrorStacktrace(err)}, fields...)...)

	return &errorReturner{err}
}

// Emergency is alias of below:
//
//	Emergency(err.Error(), rec.Error(err), rec.ErrorStacktrace(err))
func (e *errorLogger) Emergency(err error, fields ...Field) *errorReturner {
	var message string

	if errors.Is(err, nil) {
		message = formattedNilString
	} else {
		message = err.Error()
	}

	e.l.write(time.Now(), EMERGENCY, message, append([]Field{Error(err), ErrorStacktrace(err)}, fields...)...)

	return &errorReturner{err}
}
