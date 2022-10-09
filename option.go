package rec

import (
	"fmt"
	"os"
	"runtime"
)

// Option is a struct that handles the `*rec.Config` used when new `*rec.Logger`.
type Option struct {
	name string
	f    func(*Config) error
}

func funcName() (funcName string) {
	// nolint: dogsled
	pc, _, _, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()

	return name
}

// WithUseTimestampField returns `rec.Option` for setting `config.WithUseTimestampField`.
func WithUseTimestampField(use bool) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.UseTimestampField = use

			return nil
		},
	}
}

// WithTimestampFieldKey returns `rec.Option` for setting `config.WithTimestampFieldKey`.
func WithTimestampFieldKey(key string) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.TimestampFieldKey = key

			return nil
		},
	}
}

// WithTimestampFieldFormat returns `rec.Option` for setting `config.WithTimestampFieldFormat`.
func WithTimestampFieldFormat(format string) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.TimestampFieldFormat = format

			return nil
		},
	}
}

// WithUseSeverityField returns `rec.Option` for setting `config.WithUseSeverityField`.
func WithUseSeverityField(use bool) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.UseSeverityField = use

			return nil
		},
	}
}

// WithSeverityFieldKey returns `rec.Option` for setting `config.WithSeverityFieldKey`.
func WithSeverityFieldKey(key string) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.SeverityFieldKey = key

			return nil
		},
	}
}

// WithSeverityThreshold returns `rec.Option` for setting `config.WithSeverityThreshold`.
func WithSeverityThreshold(severity Severity) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.SeverityThreshold = severity

			return nil
		},
	}
}

// WithUseUppercaseSeverity returns `rec.Option` for setting `config.WithUseUppercaseSeverity`.
func WithUseUppercaseSeverity(use bool) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.UseUppercaseSeverity = use

			return nil
		},
	}
}

// WithDefaultSeverity returns `rec.Option` for setting `config.DefaultSeverity`.
func WithDefaultSeverity(severity Severity) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.DefaultSeverity = severity

			return nil
		},
	}
}

// WithUseHostnameField returns `rec.Option` for setting `config.WithUseHostnameField`.
func WithUseHostnameField(use bool) Option {
	return withUseHostnameField(use, funcName(), os.Hostname)
}

func withUseHostnameField(use bool, funcName string, osHostname func() (string, error)) Option {
	return Option{
		name: funcName,
		f: func(config *Config) error {
			hostname, err := osHostname()
			if err != nil {
				return fmt.Errorf("os.Hostname: %w", err)
			}
			config.UseHostnameField = use
			config.HostnameFieldValue = hostname

			return nil
		},
	}
}

// WithHostnameFieldKey returns `rec.Option` for setting `config.WithHostnameFieldKey`.
func WithHostnameFieldKey(key string) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.HostnameFieldKey = key

			return nil
		},
	}
}

// WithHostnameFieldValue returns `rec.Option` for setting `config.WithHostnameFieldValue`.
func WithHostnameFieldValue(value string) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.HostnameFieldValue = value

			return nil
		},
	}
}

// WithUseCallerField returns `rec.Option` for setting `config.WithUseCallerField`.
func WithUseCallerField(use bool) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.UseCallerField = use

			return nil
		},
	}
}

// WithCallerFieldKey returns `rec.Option` for setting `config.WithCallerFieldKey`.
func WithCallerFieldKey(key string) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.CallerFieldKey = key

			return nil
		},
	}
}

// WithCallerSkip returns `rec.Option` for setting `config.WithCallerSkip`.
func WithCallerSkip(callerSkip int) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.CallerSkip += callerSkip

			return nil
		},
	}
}

// WithUseShortCaller returns `rec.Option` for setting `config.WithUseShortCaller`.
func WithUseShortCaller(use bool) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.UseShortCaller = use

			return nil
		},
	}
}

// WithUseMessageField returns `rec.Option` for setting `config.WithUseMessageField`.
func WithUseMessageField(use bool) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.UseMessageField = use

			return nil
		},
	}
}

// WithMessageFieldKey returns `rec.Option` for setting `config.WithMessageFieldKey`.
func WithMessageFieldKey(key string) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.MessageFieldKey = key

			return nil
		},
	}
}

// WithLineSeparator returns `rec.Option` for setting `config.WithLineSeparator`.
func WithLineSeparator(separator string) Option {
	return Option{
		name: funcName(),
		f: func(config *Config) error {
			config.LineSeparator = separator

			return nil
		},
	}
}
