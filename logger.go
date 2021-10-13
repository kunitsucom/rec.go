package rec

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var defaultLogger = Must(NewWithConfig(os.Stderr, NewConfig()))

// Logger is the main struct of rec.
type Logger struct {
	mu sync.Mutex

	customSeverities map[Severity]*severityStrings

	config *Config

	contextFields []byte

	writer io.Writer
}

// Lock locks mu. If the lock is already in use, the calling goroutine blocks until the mutex is available.
func (l *Logger) Lock() {
	l.mu.Lock()
}

// Unlock unlocks mu. It is a run-time error if m is not locked on entry to Unlock.
// A locked Mutex is not associated with a particular goroutine. It is allowed for one goroutine to lock a Mutex and then arrange for another goroutine to unlock it.
func (l *Logger) Unlock() {
	l.mu.Unlock()
}

// NewWithConfig creates a *rec.Logger from *rec.Config.
func NewWithConfig(writer io.Writer, config *Config) (*Logger, error) {
	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("*Config.validate: %w", err)
	}

	return &Logger{
		mu:               sync.Mutex{},
		customSeverities: defaultSeverities(),
		config:           config,
		contextFields:    make([]byte, 0),
		writer:           writer,
	}, nil
}

// New creates a *rec.Logger from rec.Option.
func New(writer io.Writer, options ...Option) (*Logger, error) {
	config := NewConfig()

	for _, opt := range options {
		if err := opt.f(config); err != nil {
			return nil, fmt.Errorf("%s: %w", opt.name, err)
		}
	}

	return NewWithConfig(writer, config)
}

// Must is used as Must(rec.New()) or Must(rec.NewWithConfig()).
// If the passed function returns an error, Must() will panic.
func Must(l *Logger, err error) *Logger {
	if err != nil {
		panic(fmt.Errorf("Must: %w", err))
	}

	return l
}

// L returns default `*rec.Logger` in rec package.
func L() *Logger {
	return defaultLogger
}

// ReplaceDefaultLogger replaces the default logger in rec package.
// and returns a function for rollback logger.
func ReplaceDefaultLogger(l *Logger) (rollback func()) {
	backupDefaultLogger := defaultLogger

	defaultLogger = l

	return func() {
		defaultLogger = backupDefaultLogger
	}
}

// Copy returns copied `*rec.Logger`.
func (l *Logger) Copy() *Logger {
	copiedLogger := Must(NewWithConfig(l.writer, NewConfig()))

	// subst copied customSeverities
	for severity, severityStrings := range l.customSeverities {
		copiedSeverityStrings := *severityStrings // copy
		copiedLogger.customSeverities[severity] = &copiedSeverityStrings
	}

	// subst copied LoggerOptions
	copiedLoggerConfig := *l.config // copy
	copiedLogger.config = &copiedLoggerConfig

	// copy context fields
	if len(l.contextFields) > 0 {
		copiedLogger.contextFields = append(copiedLogger.contextFields, l.contextFields...)
	}

	return copiedLogger
}

// AddCallerSkip copied `*rec.Logger` that added caller skip.
func (l *Logger) AddCallerSkip(callerSkip int) *Logger {
	copied := l.Copy()
	copied.config.CallerSkip += callerSkip

	return copied
}

// With returns a `rec.Logger` with fixed `rec.Fields` added.
//
// For example, for the following `rec.Logger`:
//
//     l := rec.Must(rec.New(os.Stderr)).With(rec.String("field", "added"))
//     l.InfoLog("rec")
//
// This will output a log like the following:
//
//     $ go run main.go
//     {"timestamp":"...",...,"message":"rec","field":"added"}
//
func (l *Logger) With(fields ...Field) *Logger {
	copied := l.Copy()

	for i := range fields {
		copied.contextFields = append(appendJSONField(copied.contextFields, fields[i]), ',')
	}

	return copied
}

// Renew copies the `*rec.Logger`, applies `rec.Option` to it, and returns it.
func (l *Logger) Renew(options ...Option) (*Logger, error) {
	copied := l.Copy()

	for _, opt := range options {
		if err := opt.f(copied.config); err != nil {
			return nil, fmt.Errorf("%s: %w", opt.name, err)
		}
	}

	return copied, nil
}

// RenewWriter copies the `*rec.Logger`, set a new io.Writer for it, and returns it.
func (l *Logger) RenewWriter(writer io.Writer) *Logger {
	copied := l.Copy()
	copied.writer = writer

	return copied
}

// nolint: cyclop, funlen
func (l *Logger) write(now time.Time, severity Severity, message string, fields ...Field) {
	if severity < l.config.SeverityThreshold {
		return
	}

	b := bufferPool.Get().(*buffer) // nolint: forcetypeassert
	defer bufferPool.Put(b)

	// reset
	b.Buffer = b.Buffer[:0]

	// {
	b.Buffer = append(b.Buffer, '{')

	// {"timestamp":"...",
	if l.config.UseTimestampField {
		b.Buffer = append(b.Buffer, '"')
		b.Buffer = appendJSONEscapedString(b.Buffer, l.config.TimestampFieldKey)
		b.Buffer = append(b.Buffer, `":`...)
		b.Buffer = appendTimeFieldValue(b.Buffer, now, l.config.TimestampFieldFormat)
		b.Buffer = append(b.Buffer, ',')
	}

	// {"timestamp":"...","severity":"...",
	if l.config.UseSeverityField {
		b.Buffer = append(b.Buffer, '"')
		b.Buffer = appendJSONEscapedString(b.Buffer, l.config.SeverityFieldKey)
		b.Buffer = append(b.Buffer, `":"`...)

		if l.config.UseUppercaseSeverity {
			b.Buffer = appendJSONEscapedString(b.Buffer, l.uppercase(severity))
		} else {
			b.Buffer = appendJSONEscapedString(b.Buffer, l.lowercase(severity))
		}

		b.Buffer = append(b.Buffer, `",`...)
	}

	// {"timestamp":"...","severity":"...","hostname":"...",
	if l.config.UseHostnameField {
		b.Buffer = append(b.Buffer, '"')
		b.Buffer = appendJSONEscapedString(b.Buffer, l.config.HostnameFieldKey)
		b.Buffer = append(b.Buffer, `":"`...)
		b.Buffer = appendJSONEscapedString(b.Buffer, l.config.HostnameFieldValue)
		b.Buffer = append(b.Buffer, `",`...)
	}

	// {"timestamp":"...","severity":"...","hostname":"...","caller":"...",
	if l.config.UseCallerField {
		b.Buffer = append(b.Buffer, '"')
		b.Buffer = appendJSONEscapedString(b.Buffer, l.config.CallerFieldKey)
		b.Buffer = append(b.Buffer, `":"`...)
		b.Buffer = appendCaller(b.Buffer, l.config.CallerSkip, l.config.UseShortCaller)
		b.Buffer = append(b.Buffer, `",`...)
	}

	// {"timestamp":"...","severity":"...","hostname":"...","caller":"...","message":"...",
	if l.config.UseMessageField {
		b.Buffer = append(b.Buffer, '"')
		b.Buffer = appendJSONEscapedString(b.Buffer, l.config.MessageFieldKey)
		b.Buffer = append(b.Buffer, `":"`...)
		b.Buffer = appendJSONEscapedString(b.Buffer, message)
		b.Buffer = append(b.Buffer, `",`...)
	}

	// {"timestamp":"...","severity":"...","hostname":"...","caller":"...","message":"...","context":"...",
	if len(l.contextFields) > 0 {
		b.Buffer = append(b.Buffer, l.contextFields...)
	}

	// {"timestamp":"...","severity":"...","hostname":"...","caller":"...","message":"...","context":"...","fields":"...",
	for i := range fields {
		b.Buffer = append(appendJSONField(b.Buffer, fields[i]), ',')
	}

	// {"timestamp":"...","severity":"...","hostname":"...","caller":"...","message":"...","context":"...","fields":"..."}
	if b.Buffer[len(b.Buffer)-1] == ',' {
		b.Buffer[len(b.Buffer)-1] = '}'
	} else {
		b.Buffer = append(b.Buffer, '}')
	}

	// {"timestamp":"...","severity":"...","hostname":"...","caller":"...","message":"...","context":"...","fields":"..."}\n
	b.Buffer = append(b.Buffer, l.config.LineSeparator...)

	if _, err := l.writer.Write(b.Buffer); err != nil {
		err = fmt.Errorf("rec.(*Logger).write: writer=%#v: Write: %w", l.writer, err)
		defaultLogger.write(now, ERROR, err.Error(), Error(err))
	}
}
