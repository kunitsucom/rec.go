package rec

import "errors"

var (
	// ErrIsEmpty is empty.
	ErrIsEmpty = errors.New("is empty")

	// ErrSeverityAlreadyExists severity already exists.
	ErrSeverityAlreadyExists = errors.New("severity already exists")
	// ErrSeverityLowerCaseIsEmpty severity lowercase is empty.
	ErrSeverityLowerCaseIsEmpty = errors.New("severity lowercase is empty")
	// ErrSeverityUpperCaseIsEmpty severity uppercase is empty.
	ErrSeverityUpperCaseIsEmpty = errors.New("severity uppercase is empty")
)
