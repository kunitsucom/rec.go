// nolint: testpackage
package rec

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"testing"
	"time"
)

var (
	errForTest                    = errors.New("test error")
	errEnvironmentVariableIsEmpty = errors.New("environment variable is empty")
)

var formatEqual = func() string {
	if must.ParseBool(os.Getenv("COLOR")) {
		return "\n\033[31m--- noteql\033[0m\n\033[32m+++ actual\033[0m\n\033[31m-%v\033[0m\n\033[32m+%v\033[0m\n"
	}

	return "\n--- noteql\n+++ actual\n-%v\n+%v\n"
}()

var formatNotEqual = func() string {
	if must.ParseBool(os.Getenv("COLOR")) {
		return "\n\033[31m--- expect\033[0m\n\033[32m+++ actual\033[0m\n\033[31m-%v\033[0m\n\033[32m+%v\033[0m\n"
	}

	return "\n--- expect\n+++ actual\n-%v\n+%v\n"
}()

func FailIfEqual(t *testing.T, expect interface{}, actual interface{}) {
	t.Helper()

	if expect == actual {
		t.Errorf(formatEqual, expect, actual)
	}
}

func FailIfBytesEqual(t *testing.T, expect []byte, actual []byte) {
	t.Helper()

	if bytes.Equal(expect, actual) {
		t.Errorf(formatEqual, string(expect), string(actual))
	}
}

func FailIfDeepEqual(t *testing.T, expect interface{}, actual interface{}) {
	t.Helper()

	if reflect.DeepEqual(expect, actual) {
		t.Errorf(formatEqual, expect, actual)
	}
}

func FailIfErrorIs(t *testing.T, expect, actual error) {
	t.Helper()

	if errors.Is(actual, expect) {
		t.Errorf(formatEqual, expect, actual)
	}
}

func FailIfRegexpMatchString(t *testing.T, expect *regexp.Regexp, actual string) {
	t.Helper()

	if expect.MatchString(actual) {
		t.Errorf(formatEqual, expect.String(), actual)
	}
}

func FailIfNotEqual(t *testing.T, expect interface{}, actual interface{}) {
	t.Helper()

	if expect != actual {
		t.Errorf(formatNotEqual, expect, actual)
	}
}

func FailIfNotBytesEqual(t *testing.T, expect []byte, actual []byte) {
	t.Helper()

	if !bytes.Equal(expect, actual) {
		t.Errorf(formatNotEqual, string(expect), string(actual))
	}
}

func FailIfNotDeepEqual(t *testing.T, expect interface{}, actual interface{}) {
	t.Helper()

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf(formatNotEqual, expect, actual)
	}
}

func FailIfNotErrorIs(t *testing.T, expect, actual error) {
	t.Helper()

	if !errors.Is(actual, expect) {
		t.Errorf(formatNotEqual, expect, actual)
	}
}

func FailIfNotRegexpMatchString(t *testing.T, expect *regexp.Regexp, actual string) {
	t.Helper()

	if !expect.MatchString(actual) {
		t.Errorf(formatNotEqual, expect.String(), actual)
	}
}

type mustUtil struct{}

var must mustUtil

func (mustUtil) ParseBool(str string) bool {
	v, _ := strconv.ParseBool(str)

	return v
}

func (mustUtil) Getenv(env string) string {
	if v := os.Getenv(env); v != "" {
		return v
	}

	panic(errEnvironmentVariableIsEmpty)
}

func (mustUtil) OpenFile(name string, flag int, perm fs.FileMode) *os.File {
	f, err := os.OpenFile(name, flag, perm) // nolint: gosec
	if err != nil {
		panic(err)
	}

	return f
}

type ptrUtil struct{}

var ptr ptrUtil

func (ptrUtil) Bool(v bool) *bool {
	return &v
}

func (ptrUtil) Uint(v uint) *uint {
	return &v
}

func (ptrUtil) Uint8(v uint8) *uint8 {
	return &v
}

func (ptrUtil) Uint16(v uint16) *uint16 {
	return &v
}

func (ptrUtil) Uint32(v uint32) *uint32 {
	return &v
}

func (ptrUtil) Uint64(v uint64) *uint64 {
	return &v
}

func (ptrUtil) Int(v int) *int {
	return &v
}

func (ptrUtil) Int8(v int8) *int8 {
	return &v
}

func (ptrUtil) Int16(v int16) *int16 {
	return &v
}

func (ptrUtil) Int32(v int32) *int32 {
	return &v
}

func (ptrUtil) Int64(v int64) *int64 {
	return &v
}

func (ptrUtil) Float32(v float32) *float32 {
	return &v
}

func (ptrUtil) Float64(v float64) *float64 {
	return &v
}

func (ptrUtil) Complex64(v complex64) *complex64 {
	return &v
}

func (ptrUtil) Complex128(v complex128) *complex128 {
	return &v
}

func (ptrUtil) Time(v time.Time) *time.Time {
	return &v
}

func (ptrUtil) Duration(v time.Duration) *time.Duration {
	return &v
}

func (ptrUtil) String(v string) *string {
	return &v
}
