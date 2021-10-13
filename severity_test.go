// nolint: testpackage
package rec

import (
	"sync"
	"testing"
)

func TestLogger_Lowercase(t *testing.T) {
	customSeverities := defaultSeverities()
	customSeverities[50] = &severityStrings{uppercase: "TRACE", lowercase: "trace"}

	type fields struct {
		Config           *Config
		customSeverities map[Severity]*severityStrings
	}

	type args struct {
		s Severity
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		expect string
	}{
		{name: lowercaseDefault, fields: fields{Config: NewConfig(), customSeverities: customSeverities}, args: args{s: DEFAULT}, expect: lowercaseDefault},
		{name: lowercaseDebug, fields: fields{Config: NewConfig(), customSeverities: customSeverities}, args: args{s: DEBUG}, expect: lowercaseDebug},
		{name: lowercaseInfo, fields: fields{Config: NewConfig(), customSeverities: customSeverities}, args: args{s: INFO}, expect: lowercaseInfo},
		{name: lowercaseNotice, fields: fields{Config: NewConfig(), customSeverities: customSeverities}, args: args{s: NOTICE}, expect: lowercaseNotice},
		{name: lowercaseWarning, fields: fields{Config: NewConfig(), customSeverities: customSeverities}, args: args{s: WARNING}, expect: lowercaseWarning},
		{name: lowercaseError, fields: fields{Config: NewConfig(), customSeverities: customSeverities}, args: args{s: ERROR}, expect: lowercaseError},
		{name: lowercaseCritical, fields: fields{Config: NewConfig(), customSeverities: customSeverities}, args: args{s: CRITICAL}, expect: lowercaseCritical},
		{name: lowercaseAlert, fields: fields{Config: NewConfig(), customSeverities: customSeverities}, args: args{s: ALERT}, expect: lowercaseAlert},
		{name: lowercaseEmergency, fields: fields{Config: NewConfig(), customSeverities: customSeverities}, args: args{s: EMERGENCY}, expect: lowercaseEmergency},
		{name: "trace", fields: fields{Config: NewConfig(), customSeverities: customSeverities}, args: args{s: 50}, expect: "trace"},
		{name: "none", fields: fields{Config: NewConfig(), customSeverities: customSeverities}, args: args{s: 1}, expect: "1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				config:           tt.fields.Config,
				customSeverities: tt.fields.customSeverities,
			}
			actual := l.lowercase(tt.args.s)
			FailIfNotEqual(t, tt.expect, actual)
		})
	}
}

func TestLogger_Uppercase(t *testing.T) {
	customSeverities := defaultSeverities()
	customSeverities[50] = &severityStrings{uppercase: "TRACE", lowercase: "trace"}

	type fields struct {
		m                *sync.Mutex
		Options          *Config
		customSeverities map[Severity]*severityStrings
	}

	type args struct {
		s Severity
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		expect string
	}{
		{name: uppercaseDefault, fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: customSeverities}, args: args{s: DEFAULT}, expect: uppercaseDefault},
		{name: uppercaseDebug, fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: customSeverities}, args: args{s: DEBUG}, expect: uppercaseDebug},
		{name: uppercaseInfo, fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: customSeverities}, args: args{s: INFO}, expect: uppercaseInfo},
		{name: uppercaseNotice, fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: customSeverities}, args: args{s: NOTICE}, expect: uppercaseNotice},
		{name: uppercaseWarning, fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: customSeverities}, args: args{s: WARNING}, expect: uppercaseWarning},
		{name: uppercaseError, fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: customSeverities}, args: args{s: ERROR}, expect: uppercaseError},
		{name: uppercaseCritical, fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: customSeverities}, args: args{s: CRITICAL}, expect: uppercaseCritical},
		{name: uppercaseAlert, fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: customSeverities}, args: args{s: ALERT}, expect: uppercaseAlert},
		{name: uppercaseEmergency, fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: customSeverities}, args: args{s: EMERGENCY}, expect: uppercaseEmergency},
		{name: "TRACE", fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: customSeverities}, args: args{s: 50}, expect: "TRACE"},
		{name: "NONE", fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: customSeverities}, args: args{s: 1}, expect: "1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				config:           tt.fields.Options,
				customSeverities: tt.fields.customSeverities,
			}
			actual := l.uppercase(tt.args.s)
			FailIfNotEqual(t, tt.expect, actual)
		})
	}
}

func TestLogger_AddCustomSeverity(t *testing.T) {
	type fields struct {
		m                *sync.Mutex
		Options          *Config
		customSeverities map[Severity]*severityStrings
	}

	type args struct {
		severity  Severity
		lowercase string
		uppercase string
	}

	tests := []struct {
		name        string
		fields      fields
		args        args
		expectError error
	}{
		{name: "success(TRACE)", fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: defaultSeverities()}, args: args{severity: 50, lowercase: "trace", uppercase: "TRACE"}, expectError: nil},
		{name: "error(severity,ErrSeverityAlreadyExists)", fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: defaultSeverities()}, args: args{severity: DEFAULT, lowercase: "not-exists", uppercase: "not-exists"}, expectError: ErrSeverityAlreadyExists},
		{name: "error(lowercase,ErrSeverityAlreadyExists)", fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: defaultSeverities()}, args: args{severity: 1, lowercase: lowercaseDefault, uppercase: "not-exists"}, expectError: ErrSeverityAlreadyExists},
		{name: "error(uppercase,ErrSeverityAlreadyExists)", fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: defaultSeverities()}, args: args{severity: 1, lowercase: "not-exists", uppercase: uppercaseDefault}, expectError: ErrSeverityAlreadyExists},
		{name: "error(ErrSeverityLowerCaseIsEmpty)", fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: defaultSeverities()}, args: args{severity: 1, lowercase: "", uppercase: "not-exists"}, expectError: ErrSeverityLowerCaseIsEmpty},
		{name: "error(ErrSeverityUpperCaseIsEmpty)", fields: fields{m: &sync.Mutex{}, Options: NewConfig(), customSeverities: defaultSeverities()}, args: args{severity: 1, lowercase: "not-exists", uppercase: ""}, expectError: ErrSeverityUpperCaseIsEmpty},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				config:           tt.fields.Options,
				customSeverities: tt.fields.customSeverities,
			}
			actual := l.AddCustomSeverity(tt.args.severity, tt.args.lowercase, tt.args.uppercase)
			FailIfNotErrorIs(t, tt.expectError, actual)
		})
	}
}
