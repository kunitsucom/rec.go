// nolint: testpackage
package rec

import (
	"os"
	"regexp"
	"testing"
)

func Test_funcName(t *testing.T) {
	tests := []struct {
		name           string
		expectFuncName string
	}{
		{"Test_funcName", "Test_funcName"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expect := regexp.MustCompile(`.*Test_funcName.*`)
			actual := funcName()
			FailIfNotRegexpMatchString(t, expect, actual)
		})
	}
}

func TestUseSeverityField(t *testing.T) {
	tests := []struct {
		name   string
		use    bool
		expect error
	}{
		{"success()", false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithUseSeverityField(tt.use)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.use, config.UseSeverityField)
		})
	}
}

func TestSeverityFieldKey(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		expect error
	}{
		{"success()", "test", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithSeverityFieldKey(tt.key)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.key, config.SeverityFieldKey)
		})
	}
}

func TestSeverityThreshold(t *testing.T) {
	tests := []struct {
		name     string
		severity Severity
		expect   error
	}{
		{"success()", 1, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithSeverityThreshold(tt.severity)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.severity, config.SeverityThreshold)
		})
	}
}

func TestUseUppercaseSeverity(t *testing.T) {
	tests := []struct {
		name   string
		use    bool
		expect error
	}{
		{"success()", false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithUseUppercaseSeverity(tt.use)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.use, config.UseHostnameField)
		})
	}
}

func TestUseTimestampField(t *testing.T) {
	tests := []struct {
		name   string
		use    bool
		expect error
	}{
		{"success()", false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithUseTimestampField(tt.use)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.use, config.UseHostnameField)
		})
	}
}

func TestTimestampFieldKey(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		expect error
	}{
		{"success()", "test", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithTimestampFieldKey(tt.key)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.key, config.TimestampFieldKey)
		})
	}
}

func TestTimestampFieldFormat(t *testing.T) {
	tests := []struct {
		name   string
		format string
		expect error
	}{
		{"success()", "1504", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithTimestampFieldFormat(tt.format)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.format, config.TimestampFieldFormat)
		})
	}
}

func TestUseHostnameField(t *testing.T) {
	tests := []struct {
		name   string
		use    bool
		expect error
	}{
		{"success()", true, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithUseHostnameField(tt.use)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.use, config.UseHostnameField)
		})
	}
}

func Test_useHostnameField(t *testing.T) {
	tests := []struct {
		name       string
		use        bool
		funcName   string
		osHostname func() (string, error)
		expect     error
	}{
		{"success()", true, "test", os.Hostname, nil},
		{"error()", false, "test", func() (string, error) { return "", errForTest }, errForTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := withUseHostnameField(tt.use, tt.funcName, tt.osHostname)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.use, config.UseHostnameField)
		})
	}
}

func TestHostnameFieldKey(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		expect error
	}{
		{"success()", "test", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithHostnameFieldKey(tt.key)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.key, config.HostnameFieldKey)
		})
	}
}

func TestHostnameFieldValue(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		expect error
	}{
		{"success()", "test", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithHostnameFieldValue(tt.value)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.value, config.HostnameFieldValue)
		})
	}
}

func TestUseCallerField(t *testing.T) {
	tests := []struct {
		name   string
		use    bool
		expect error
	}{
		{"success()", false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithUseCallerField(tt.use)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.use, config.UseCallerField)
		})
	}
}

func TestCallerFieldKey(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		expect error
	}{
		{"success()", "test", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithCallerFieldKey(tt.key)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.key, config.CallerFieldKey)
		})
	}
}

func TestCallerSkip(t *testing.T) {
	tests := []struct {
		name   string
		skip   int
		expect error
	}{
		{"success()", 4, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithCallerSkip(tt.skip)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.skip+defaultCallerSkip, config.CallerSkip)
		})
	}
}

func TestUseShortCaller(t *testing.T) {
	tests := []struct {
		name   string
		use    bool
		expect error
	}{
		{"success()", false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithUseShortCaller(tt.use)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.use, config.UseShortCaller)
		})
	}
}

func TestUseMessageField(t *testing.T) {
	tests := []struct {
		name   string
		use    bool
		expect error
	}{
		{"success()", false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithUseMessageField(tt.use)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.use, config.UseMessageField)
		})
	}
}

func TestMessageFieldKey(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		expect error
	}{
		{"success()", "test", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithMessageFieldKey(tt.key)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.key, config.MessageFieldKey)
		})
	}
}

func TestLineSeparator(t *testing.T) {
	tests := []struct {
		name      string
		separator string
		expect    error
	}{
		{"success()", "\r\n", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			option := WithLineSeparator(tt.separator)
			actual := option.f(config)
			FailIfNotErrorIs(t, tt.expect, actual)
			FailIfNotEqual(t, tt.separator, config.LineSeparator)
		})
	}
}
