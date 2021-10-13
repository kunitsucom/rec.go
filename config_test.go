// nolint: testpackage
package rec

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Parallel()

	t.Run("success(NewConfig)", func(t *testing.T) {
		t.Parallel()

		expect, actual := NewConfig(), newConfig(os.Hostname)
		FailIfNotDeepEqual(t, expect, actual)
	})
}

func Test_newConfig(t *testing.T) {
	t.Parallel()

	errFunc := func() (string, error) {
		return "", errForTest
	}
	errConfig := newConfig(os.Hostname)
	errConfig.HostnameFieldValue = "localhost"

	type args struct {
		osHostname func() (string, error)
	}

	tests := []struct {
		name   string
		args   args
		expect *Config
	}{
		{"success(newConfig)", args{os.Hostname}, newConfig(os.Hostname)},
		{"error(os.Hostname)", args{errFunc}, errConfig},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			expect, actual := tt.expect, newConfig(tt.args.osHostname)
			FailIfNotDeepEqual(t, expect, actual)
		})
	}
}

func TestConfig_validate(t *testing.T) {
	t.Parallel()

	configOK := NewConfig()

	configOKTimestampFieldKey := NewConfig()
	configOKTimestampFieldKey.UseTimestampField = false
	configOKTimestampFieldKey.TimestampFieldKey = ""

	configNGTimestampFieldKey := NewConfig()
	configNGTimestampFieldKey.TimestampFieldKey = ""

	configOKSeverityFieldKey := NewConfig()
	configOKSeverityFieldKey.UseSeverityField = false
	configOKSeverityFieldKey.SeverityFieldKey = ""

	configNGSeverityFieldKey := NewConfig()
	configNGSeverityFieldKey.SeverityFieldKey = ""

	configOKCallerFieldKey := NewConfig()
	configOKCallerFieldKey.UseCallerField = false
	configOKCallerFieldKey.CallerFieldKey = ""

	configNGCallerFieldKey := NewConfig()
	configNGCallerFieldKey.CallerFieldKey = ""

	configOKMessageFieldKey := NewConfig()
	configOKMessageFieldKey.UseMessageField = false
	configOKMessageFieldKey.MessageFieldKey = ""

	configNGMessageFieldKey := NewConfig()
	configNGMessageFieldKey.MessageFieldKey = ""

	tests := []struct {
		name      string
		config    *Config
		expectErr error
	}{
		{"success(validate)", configOK, nil},
		{"success(TimestampFieldKey)", configOKTimestampFieldKey, nil},
		{"error(TimestampFieldKey)", configNGTimestampFieldKey, ErrIsEmpty},
		{"success(SeverityFieldKey)", configOKSeverityFieldKey, nil},
		{"error(SeverityFieldKey)", configNGSeverityFieldKey, ErrIsEmpty},
		{"success(CallerFieldKey)", configOKCallerFieldKey, nil},
		{"error(CallerFieldKey)", configNGCallerFieldKey, ErrIsEmpty},
		{"success(MessageFieldKey)", configOKMessageFieldKey, nil},
		{"error(MessageFieldKey)", configNGMessageFieldKey, ErrIsEmpty},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			expect, actual := tt.expectErr, tt.config.validate()
			FailIfNotErrorIs(t, expect, actual)
		})
	}
}
