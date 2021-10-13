// nolint: testpackage
package rec

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"testing"
	"time"
)

const (
	testLogEntryMessage           = "\"🚀🚚🎉🙏\x01\b\f\n\r\t🚀🚚🎉🙏\""
	testLogEntryMessageJSONEscape = `\"🚀🚚🎉🙏\u0001\b\f\n\r\t🚀🚚🎉🙏\"`
)

// nolint: deadcode, unused, varcheck
var (
	devnull, _                     = os.OpenFile(os.DevNull, os.O_RDWR, 0o600)
	tzJST, _                       = time.LoadLocation("Asia/Tokyo")
	testTimestampValue             = time.Unix(1609464225, 678900000).In(tzJST) // 2021-01-01T10:23:45.6789+09:00
	testBoolKey                    = "testBoolKey"
	testBoolValueTrue              = true
	testBoolValueFalse             = false
	testBoolPtrKey                 = "testBoolKey"
	testUintKey                    = "testUintKey"
	testUintValue        uint      = math.MaxUint
	testUintPtrKey                 = "testUintPtrKey"
	testUint8Key                   = "testUint8Key"
	testUint8Value       uint8     = math.MaxUint8
	testUint8PtrKey                = "testUint8PtrKey"
	testUint16Key                  = "testUint16Key"
	testUint16Value      uint16    = math.MaxUint16
	testUint16PtrKey               = "testUint16PtrKey"
	testUint32Key                  = "testUint32Key"
	testUint32Value      uint32    = math.MaxUint32
	testUint64Key                  = "testUint64Key"
	testUint64Value      uint64    = math.MaxUint64
	testIntKey                     = "testIntKey"
	testIntValue                   = math.MinInt
	testInt8Key                    = "testInt8Key"
	testInt8Value        int8      = math.MaxInt8
	testInt16Key                   = "testInt16Key"
	testInt16Value       int16     = math.MaxInt16
	testInt32Key                   = "testInt32Key"
	testInt32Value       int32     = math.MaxInt32
	testInt64Key                   = "testInt64Key"
	testInt64Value       int64     = math.MaxInt64
	testFloat32Key                 = "testFloat32Key"
	testFloat32Value     float32   = math.Pi
	testComplex64Key               = "testComplex64Key"
	testComplex64Value   complex64 = complex(math.MaxFloat32, math.MaxFloat32)
	testTimeKey                    = "testTimeKey"
	testTimeValue                  = time.Unix(1609464225, 678900000).In(time.Local)
	testTimePtrKey                 = "testTimePtrKey"
	testTimeFormatKey              = "testTimeFormatKey"
	testTimeFormatPtrKey           = "testTimeFormatPtrKey"
	testDurationKey                = "testDurationKey"
	testDurationValue              = time.Minute
	testStringKey                  = "testStringKey"
	testStringValue                = "testStringValue"
	testStringsKey                 = "testStringsKey"
	testStringsValue               = []string{"test string 0", "test string 1"}
	testStringPtrKey               = "testStringPtrKey"
	testStringPtrValue             = "testStringPtrValue"
	testStringerKey                = "testStringerKey"
	testStringerValue              = time.Date(1919, 8, 18, 8, 9, 3, 114514364, time.UTC)
	testErrorValue                 = errors.New("test error") // nolint: errname, revive, stylecheck
	testErrorsValue                = []error{fmt.Errorf("%w 0", testErrorValue), fmt.Errorf("%w 1", testErrorValue)}
)

func TestNewWithConfig(t *testing.T) {
	t.Parallel()

	configOK := NewConfig()

	configNG := NewConfig()
	configNG.SeverityFieldKey = ""

	tests := []struct {
		name        string
		config      *Config
		expectError error
	}{
		{"success(NewWithConfig)", configOK, nil},
		{"error(validate)", configNG, ErrIsEmpty},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			expect := tt.expectError
			_, actual := NewWithConfig(io.Discard, tt.config)
			FailIfNotErrorIs(t, expect, actual)
		})
	}
}

func TestNew(t *testing.T) {
	t.Parallel()

	funcOK := Option{
		name: "ok",
		f: func(*Config) error {
			return nil
		},
	}
	funcNG := Option{
		name: "ok",
		f: func(*Config) error {
			return errForTest
		},
	}

	tests := []struct {
		name    string
		options []Option
		expect  error
	}{
		{"success(New)", []Option{funcOK}, nil},
		{"success(Option)", []Option{funcNG}, errForTest},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			expect := tt.expect
			_, actual := New(io.Discard, tt.options...)
			FailIfNotErrorIs(t, expect, actual)
		})
	}
}

func TestMust(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		l      *Logger
		err    error
		expect error
	}{
		{"success(Must)", nil, nil, nil},
		{"panic(Must)", nil, errForTest, errForTest},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			expect := tt.expect
			defer func() {
				if err := recover(); err != nil {
					actual, ok := err.(error)
					if !ok {
						return
					}
					FailIfNotErrorIs(t, expect, actual)
				}
			}()
			Must(tt.l, tt.err)
		})
	}
}

func TestLogger_Copy(t *testing.T) {
	t.Parallel()

	t.Run("success(Copy)", func(t *testing.T) {
		t.Parallel()

		original := Must(New(io.Discard))
		copied := original.Copy()

		FailIfEqual(t, fmt.Sprintf("%p", original), fmt.Sprintf("%p", copied))
		// customSeverities
		FailIfEqual(t, fmt.Sprintf("%p", original.customSeverities), fmt.Sprintf("%p", copied.customSeverities))
		FailIfNotDeepEqual(t, original.customSeverities, copied.customSeverities)
		// contextFields
		// FailIfEqual(t, fmt.Sprintf("%p", original.contextFields), fmt.Sprintf("%p", copied.contextFields)) // NOTE: Two distinct zero-size variables may have the same address in memory. https://go.dev/ref/spec#Size_and_alignment_guarantees
		FailIfNotDeepEqual(t, original.contextFields, copied.contextFields)
	})

	t.Run("success(Copy)", func(t *testing.T) {
		t.Parallel()

		original := Must(New(io.Discard)).With(String("field", "context"))
		copied := original.Copy()

		FailIfEqual(t, fmt.Sprintf("%p", original), fmt.Sprintf("%p", copied))
		// customSeverities
		FailIfEqual(t, fmt.Sprintf("%p", original.customSeverities), fmt.Sprintf("%p", copied.customSeverities))
		FailIfNotDeepEqual(t, original.customSeverities, copied.customSeverities)
		// contextFields
		FailIfEqual(t, fmt.Sprintf("%p", original.contextFields), fmt.Sprintf("%p", copied.contextFields))
		FailIfNotDeepEqual(t, original.contextFields, copied.contextFields)
	})
}

func TestLogger_AddCallerSkip(t *testing.T) {
	t.Parallel()

	t.Run("success(AddCallerSkip)", func(t *testing.T) {
		t.Parallel()

		l := Must(New(io.Discard))
		const delta = 1
		expect := l.config.CallerSkip + delta
		actual := l.AddCallerSkip(delta).config.CallerSkip

		FailIfNotEqual(t, expect, actual)
	})
}

func TestLogger_AddContextFields(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		message       string
		contextFields []Field
		field         []Field
		expect        string
	}{
		{"success()", "test", []Field{String("field", "context")}, []Field{String("field", "onetime")}, `{"severity":"INFO","message":"test","field":"context","field":"onetime"}` + defaultLineSeparator},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			buf := bytes.NewBuffer(nil)
			l := Must(New(buf, WithUseTimestampField(false), WithUseCallerField(false)))

			contextLogger := l.With(tt.contextFields...)
			contextLogger.Info(tt.message, tt.field...)
			actual := buf.String()
			FailIfNotEqual(t, tt.expect, actual)
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	funcOK := Option{
		name: "ok",
		f: func(*Config) error {
			return nil
		},
	}
	funcNG := Option{
		name: "ok",
		f: func(*Config) error {
			return errForTest
		},
	}

	tests := []struct {
		name    string
		options []Option
		expect  error
	}{
		{"success(nil)", nil, nil},
		{"success(New)", []Option{funcOK}, nil},
		{"success(Option)", []Option{funcNG}, errForTest},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			l := Must(NewWithConfig(io.Discard, NewConfig()))
			expect := tt.expect
			_, actual := l.Renew(tt.options...)
			FailIfNotErrorIs(t, expect, actual)
		})
	}
}

func TestLogger_UpdateWriter(t *testing.T) {
	t.Parallel()

	buf := bytes.NewBuffer(nil)
	tests := []struct {
		name         string
		writer       io.Writer
		expectWriter io.Writer
	}{
		{"success()", buf, buf},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			l := Must(New(io.Discard))
			actual := l.RenewWriter(buf)
			FailIfNotEqual(t, tt.expectWriter, actual.writer)
		})
	}
}

func TestLogger_print(t *testing.T) {
	t.Parallel()

	t.Run("success(severity<l.config.SeverityThreshold)", func(t *testing.T) {
		t.Parallel()

		// prepare
		l := Must(New(devnull, WithSeverityThreshold(1)))
		// run
		l.write(time.Now(), DEFAULT, testLogEntryMessage)
	})

	t.Run("success(UppercaseSeverity)", func(t *testing.T) {
		t.Parallel()

		// prepare
		l := Must(New(devnull))
		// run
		l.write(time.Now(), DEFAULT, testLogEntryMessage)
	})

	t.Run("success(LowercaseSeverity)", func(t *testing.T) {
		t.Parallel()

		// prepare
		l := Must(New(devnull, WithUseUppercaseSeverity(false)))
		// run
		l.write(time.Now(), DEFAULT, testLogEntryMessage)
	})

	t.Run("success(UseShortCaller)", func(t *testing.T) {
		t.Parallel()

		// prepare
		l := Must(New(devnull))
		// run
		l.write(time.Now(), DEFAULT, testLogEntryMessage)
	})

	t.Run("success(UseLongCaller)", func(t *testing.T) {
		t.Parallel()

		// prepare
		l := Must(New(devnull, WithUseShortCaller(false)))
		// run
		l.write(time.Now(), DEFAULT, testLogEntryMessage)
	})

	t.Run("success(UseHostnameField=true)", func(t *testing.T) {
		t.Parallel()

		// prepare
		l := Must(New(devnull, WithUseHostnameField(true)))
		// run
		l.write(time.Now(), DEFAULT, testLogEntryMessage)
	})

	t.Run("success(UseHostnameField=false)", func(t *testing.T) {
		t.Parallel()

		// prepare
		l := Must(New(devnull))
		// run
		l.write(time.Now(), DEFAULT, testLogEntryMessage)
	})

	t.Run("success(NoField)", func(t *testing.T) {
		t.Parallel()

		// prepare
		buf := bytes.NewBuffer(nil)
		l := Must(New(buf))
		l.config = &Config{}
		// run
		l.write(time.Now(), DEFAULT, testLogEntryMessage)
		// check
		const expect = `{}`
		actual := buf.String()
		FailIfNotEqual(t, expect, actual)
	})

	t.Run("error(typeNone)", func(t *testing.T) {
		t.Parallel()

		// prepare
		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseCallerField(false)))
		// run
		l.write(testTimestampValue, DEFAULT, testLogEntryMessage, Field{key: "noneField"})
		// check
		const expect = `{"timestamp":"2021-01-01T10:23:45.6789+09:00","severity":"DEFAULT","message":"` + testLogEntryMessageJSONEscape + `","noneField":"ERROR: TYPE NONE"}` + defaultLineSeparator
		actualErr := buf.String()
		FailIfNotEqual(t, expect, actualErr)
	})

	t.Run("error(undefinedField)", func(t *testing.T) {
		t.Parallel()

		// prepare
		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseCallerField(false)))
		// run
		l.write(testTimestampValue, DEFAULT, testLogEntryMessage, Field{key: "undefinedField", t: math.MaxUint8})
		// check
		const expect = `{"timestamp":"2021-01-01T10:23:45.6789+09:00","severity":"DEFAULT","message":"` + testLogEntryMessageJSONEscape + `","undefinedField":"ERROR: UNDEFINED TYPE: 255"}` + defaultLineSeparator
		actual := buf.String()
		FailIfNotEqual(t, expect, actual)
	})

	t.Run("success(OutputMatch)", func(t *testing.T) {
		t.Parallel()

		// prepare
		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithTimestampFieldFormat(""), WithUseHostnameField(true)))
		// run
		l.write(testTimestampValue, INFO, testLogEntryMessage)
		// check
		expect := regexp.MustCompile(`^{"timestamp":1609464225.6789,"severity":"INFO","hostname":".+","caller":"[^"]+:[0-9]+","message":".+"}` + l.config.LineSeparator)
		actual := buf.String()
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("error(Write)", func(t *testing.T) {
		t.Parallel()

		// prepare
		buf := bytes.NewBuffer(nil)
		replacer := Must(NewWithConfig(buf, NewConfig()))
		replacer.config.UseCallerField = false

		backup := defaultLogger
		t.Cleanup(func() { defaultLogger = backup })
		defaultLogger = replacer
		// run
		noSuchFile, _ := os.OpenFile("/tmp/no/such/file", os.O_RDWR, 0o600)
		l := Must(New(noSuchFile))
		l.write(testTimestampValue, DEFAULT, testLogEntryMessage)
		// check
		const expect = `{"timestamp":"2021-01-01T10:23:45.6789+09:00","severity":"ERROR","message":"rec.(*Logger).write: writer=(*os.File)(nil): Write: invalid argument","error":"rec.(*Logger).write: writer=(*os.File)(nil): Write: invalid argument"}` + defaultLineSeparator
		actual := buf.String()
		FailIfNotEqual(t, expect, actual)
	})
}

// nolint: paralleltest
func TestL(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	replacer := Must(New(buf))

	backup := defaultLogger

	t.Cleanup(func() { defaultLogger = backup }) // nolint: paralleltest

	defaultLogger = replacer // nolint: paralleltest

	t.Run("success(Print)", func(t *testing.T) {
		t.Cleanup(buf.Reset)
		L().Print(DEFAULT, "test")
		expect := regexp.MustCompile(`^{"timestamp":"[0-9T:\.\+\-Z]+","severity":"DEFAULT","caller":"[^"]+:[0-9]+","message":"test"}` + defaultLineSeparator)
		actual := buf.String()
		FailIfNotRegexpMatchString(t, expect, actual)
	})
}

func TestReplaceDefaultLogger(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    *Logger
	}{
		{"success()", Must(New(os.Stderr))},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			backup := defaultLogger
			FailIfEqual(t, fmt.Sprintf("%p", defaultLogger), fmt.Sprintf("%p", tt.l))
			FailIfNotEqual(t, fmt.Sprintf("%p", defaultLogger), fmt.Sprintf("%p", backup))

			rollback := ReplaceDefaultLogger(tt.l)
			FailIfNotEqual(t, fmt.Sprintf("%p", defaultLogger), fmt.Sprintf("%p", tt.l))
			FailIfEqual(t, fmt.Sprintf("%p", defaultLogger), fmt.Sprintf("%p", backup))

			rollback()
			FailIfEqual(t, fmt.Sprintf("%p", defaultLogger), fmt.Sprintf("%p", tt.l))
			FailIfNotEqual(t, fmt.Sprintf("%p", defaultLogger), fmt.Sprintf("%p", backup))
		})
	}
}
