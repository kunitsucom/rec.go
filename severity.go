package rec

import (
	"fmt"
	"strconv"
)

// Severity controls the severity of the log output by `*rec.Logger`.
// By default, the following severities are available: `DEFAULT`, `DEBUG`, `INFO`, `NOTICE`, `WARNING`, `ERROR`, `CRITICAL`, `ALERT`, and `EMERGENCY`.
type Severity int

type severityStrings struct {
	lowercase string
	uppercase string
}

// cf. https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#logseverity
const (
	DEFAULT   Severity = iota
	DEBUG     Severity = 100
	INFO      Severity = 200
	NOTICE    Severity = 300
	WARNING   Severity = 400
	ERROR     Severity = 500
	CRITICAL  Severity = 600
	ALERT     Severity = 700
	EMERGENCY Severity = 800
)

const (
	lowercaseDefault   = "default"
	lowercaseDebug     = "debug"
	lowercaseInfo      = "info"
	lowercaseNotice    = "notice"
	lowercaseWarning   = "warning"
	lowercaseError     = "error"
	lowercaseCritical  = "critical"
	lowercaseAlert     = "alert"
	lowercaseEmergency = "emergency"
)

const (
	uppercaseDefault   = "DEFAULT"
	uppercaseDebug     = "DEBUG"
	uppercaseInfo      = "INFO"
	uppercaseNotice    = "NOTICE"
	uppercaseWarning   = "WARNING"
	uppercaseError     = "ERROR"
	uppercaseCritical  = "CRITICAL"
	uppercaseAlert     = "ALERT"
	uppercaseEmergency = "EMERGENCY"
)

// lowercase returns the lowercase string of Severity.
// nolint: cyclop
func (l *Logger) lowercase(severity Severity) string {
	switch severity {
	case DEFAULT:
		return lowercaseDefault
	case DEBUG:
		return lowercaseDebug
	case INFO:
		return lowercaseInfo
	case NOTICE:
		return lowercaseNotice
	case WARNING:
		return lowercaseWarning
	case ERROR:
		return lowercaseError
	case CRITICAL:
		return lowercaseCritical
	case ALERT:
		return lowercaseAlert
	case EMERGENCY:
		return lowercaseEmergency
	default:
		if l != nil && l.customSeverities != nil && l.customSeverities[severity] != nil {
			return l.customSeverities[severity].lowercase
		}

		return strconv.Itoa(int(severity))
	}
}

// uppercase returns the uppercase string of Severity.
// nolint: cyclop
func (l *Logger) uppercase(severity Severity) string {
	switch severity {
	case DEFAULT:
		return uppercaseDefault
	case DEBUG:
		return uppercaseDebug
	case INFO:
		return uppercaseInfo
	case NOTICE:
		return uppercaseNotice
	case WARNING:
		return uppercaseWarning
	case ERROR:
		return uppercaseError
	case CRITICAL:
		return uppercaseCritical
	case ALERT:
		return uppercaseAlert
	case EMERGENCY:
		return uppercaseEmergency
	default:
		if l != nil && l.customSeverities != nil && l.customSeverities[severity] != nil {
			return l.customSeverities[severity].uppercase
		}

		return strconv.Itoa(int(severity))
	}
}

func defaultSeverities() map[Severity]*severityStrings {
	return map[Severity]*severityStrings{
		DEFAULT:   {lowercase: lowercaseDefault, uppercase: uppercaseDefault},
		DEBUG:     {lowercase: lowercaseDebug, uppercase: uppercaseDebug},
		INFO:      {lowercase: lowercaseInfo, uppercase: uppercaseInfo},
		NOTICE:    {lowercase: lowercaseNotice, uppercase: uppercaseNotice},
		WARNING:   {lowercase: lowercaseWarning, uppercase: uppercaseWarning},
		ERROR:     {lowercase: lowercaseError, uppercase: uppercaseError},
		CRITICAL:  {lowercase: lowercaseCritical, uppercase: uppercaseCritical},
		ALERT:     {lowercase: lowercaseAlert, uppercase: uppercaseAlert},
		EMERGENCY: {lowercase: lowercaseEmergency, uppercase: uppercaseEmergency},
	}
}

// AddCustomSeverity is a method to add your own Severity to the `*rec.Logger`.
func (l *Logger) AddCustomSeverity(severity Severity, lowercase, uppercase string) error {
	l.Lock()
	defer l.Unlock()

	for customSeverity, severityStrings := range l.customSeverities {
		if severity == customSeverity {
			return fmt.Errorf("severity=%d: %w", severity, ErrSeverityAlreadyExists)
		}

		if lowercase == severityStrings.lowercase {
			return fmt.Errorf("lowercase=%s: %w", lowercase, ErrSeverityAlreadyExists)
		}

		if uppercase == severityStrings.uppercase {
			return fmt.Errorf("uppercase=%s: %w", uppercase, ErrSeverityAlreadyExists)
		}

		if lowercase == "" {
			return ErrSeverityLowerCaseIsEmpty
		}

		if uppercase == "" {
			return ErrSeverityUpperCaseIsEmpty
		}
	}

	l.customSeverities[severity] = &severityStrings{
		uppercase: uppercase,
		lowercase: lowercase,
	}

	return nil
}
