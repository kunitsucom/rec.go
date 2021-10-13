package rec

import (
	"fmt"
	"os"
	"time"
)

const (
	defaultHostnameFieldValue = "localhost"
	defaultCallerSkip         = 4
	defaultLineSeparator      = "\n"
)

// Config is configuration struct for *rec.Logger.
type Config struct {
	// [timestamp] Set true if you want to output the timestamp field in the log.
	UseTimestampField bool
	// [timestamp] Set the key name in the timestamp field.
	TimestampFieldKey string
	// [timestamp] Set the Go time format for timestamp field.
	TimestampFieldFormat string

	// [severity] Set true if you want to output the severity field in the log.
	UseSeverityField bool
	// [severity] Set the key name in the severity field.
	SeverityFieldKey string
	// [severity] Set the severity of the log output.
	SeverityThreshold Severity
	// [severity] Set true if you want to output severity in uppercase.
	UseUppercaseSeverity bool

	// [hostname] Set true if you want to output the hostname field in the log.
	UseHostnameField bool
	// [hostname] Set the key name in the hostname field.
	HostnameFieldKey string
	// [hostname] Set the value in the hostname field.
	HostnameFieldValue string

	// [caller]
	UseCallerField bool
	// [caller]
	CallerFieldKey string
	// [caller]
	CallerSkip int
	// [caller]
	UseShortCaller bool

	// [message]
	UseMessageField bool
	// [message]
	MessageFieldKey string

	// [lineseparator]
	LineSeparator string
}

// NewConfig returns *rec.Config that set default values.
func NewConfig() *Config {
	return newConfig(os.Hostname)
}

func newConfig(osHostname func() (string, error)) *Config {
	config := &Config{
		// "timestamp":"...",
		UseTimestampField:    true,
		TimestampFieldKey:    "timestamp",
		TimestampFieldFormat: time.RFC3339Nano,
		// "severity":"...",
		UseSeverityField:     true,
		SeverityFieldKey:     "severity",
		SeverityThreshold:    DEFAULT,
		UseUppercaseSeverity: true,
		// "hostname":"...",
		UseHostnameField:   false,
		HostnameFieldKey:   "hostname",
		HostnameFieldValue: defaultHostnameFieldValue,
		// "caller":"...",
		UseCallerField: true,
		CallerFieldKey: "caller",
		CallerSkip:     defaultCallerSkip,
		UseShortCaller: true,
		// "message":"...",
		UseMessageField: true,
		MessageFieldKey: "message",
		// \n
		LineSeparator: defaultLineSeparator,
	}

	var err error

	config.HostnameFieldValue, err = osHostname()
	if err != nil {
		config.HostnameFieldValue = defaultHostnameFieldValue
	}

	return config
}

func (c *Config) validate() error {
	if c.UseSeverityField && c.SeverityFieldKey == "" {
		return fmt.Errorf("*Config.SeverityFieldKey %w", ErrIsEmpty)
	}

	if c.UseTimestampField && c.TimestampFieldKey == "" {
		return fmt.Errorf("*Config.TimestampFieldKey %w", ErrIsEmpty)
	}

	if c.UseCallerField && c.CallerFieldKey == "" {
		return fmt.Errorf("*Config.CallerFieldKey %w", ErrIsEmpty)
	}

	if c.UseMessageField && c.MessageFieldKey == "" {
		return fmt.Errorf("*Config.MessageFieldKey %w", ErrIsEmpty)
	}

	return nil
}
