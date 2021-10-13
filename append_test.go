// nolint: testpackage
package rec

import (
	"math"
	"runtime"
	"testing"
	"time"
)

func Test_appendStringJSONEscape(t *testing.T) {
	t.Parallel()

	t.Run(`success(🙏🙇🎉📝🚀)`, func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)
		expect := []byte(`🙏🙇🎉📝🚀`)
		actual := appendJSONEscapedString(bs, `🙏🙇🎉📝🚀`)
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run(`success(")`, func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)
		expect := []byte(`\"`)
		actual := appendJSONEscapedString(bs, "\"")
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run(`success(\)`, func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)
		expect := []byte(`\\`)
		actual := appendJSONEscapedString(bs, "\\")
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run(`success(\b)`, func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)
		expect := []byte(`\b`)
		actual := appendJSONEscapedString(bs, "\b")
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run(`success(\f)`, func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)
		expect := []byte(`\f`)
		actual := appendJSONEscapedString(bs, "\f")
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run(`success(\n)`, func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)
		expect := []byte(`\n`)
		actual := appendJSONEscapedString(bs, "\n")
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run(`success(\r)`, func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)
		expect := []byte(`\r`)
		actual := appendJSONEscapedString(bs, "\r")
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run(`success(\t)`, func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)
		expect := []byte(`\t`)
		actual := appendJSONEscapedString(bs, "\t")
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run(`success(\u0000)`, func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)
		expect := []byte(`\u0000`)
		actual := appendJSONEscapedString(bs, "\x00")
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run(`success(\u001f)`, func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)
		expect := []byte(`\u001f`)
		actual := appendJSONEscapedString(bs, "\x1f")
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run(`success(\x20)`, func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)
		expect := []byte{0x0020}
		actual := appendJSONEscapedString(bs, "\x20")
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run(`success(\xff)`, func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)
		expect := []byte{0x00ff}
		actual := appendJSONEscapedString(bs, "\xff")
		FailIfNotBytesEqual(t, expect, actual)
	})
}

func Test_appendFloatFieldValue(t *testing.T) {
	t.Parallel()

	t.Run("success(NaN)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"NaN"`)
		actual := appendFloatFieldValue(bs, math.NaN(), 64)
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(+Inf)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"+Inf"`)
		actual := appendFloatFieldValue(bs, math.Inf(1), 64)
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(-Inf)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"-Inf"`)
		actual := appendFloatFieldValue(bs, -math.Inf(1), 64)
		FailIfNotBytesEqual(t, expect, actual)
	})
}

func Test_appendCallerFromFrame(t *testing.T) {
	t.Parallel()

	t.Run("success(useShortCaller=false)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte("/a/b/c:10")
		actual := appendCallerFromFrame(bs, runtime.Frame{File: "/a/b/c", Line: 10}, false)
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(useShortCaller=true)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte("b/c:10")
		actual := appendCallerFromFrame(bs, runtime.Frame{File: "/a/b/c", Line: 10}, true)
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(abnormal,1)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(":0")
		actual := appendCallerFromFrame(bs, runtime.Frame{}, true)
		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(abnormal,2)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte("/a:10")
		actual := appendCallerFromFrame(bs, runtime.Frame{File: "/a", Line: 10}, true)
		FailIfNotBytesEqual(t, expect, actual)
	})
}

func Test_appendTimeFieldValue(t *testing.T) {
	t.Parallel()

	type args struct {
		dst    []byte
		t      time.Time
		format string
	}

	tests := []struct {
		name   string
		args   args
		expect []byte
	}{
		{name: time.RFC3339Nano, args: args{dst: make([]byte, 0), t: time.Unix(0, 0).UTC(), format: time.RFC3339Nano}, expect: []byte(`"1970-01-01T00:00:00Z"`)},
		{name: TimeFormatUnixDecimal + `(format="")`, args: args{dst: make([]byte, 0), t: time.Unix(0, 0).UTC(), format: ""}, expect: []byte(`0`)},
		{name: TimeFormatUnixDecimal + `(999ms)`, args: args{dst: make([]byte, 0), t: time.Unix(0, 999000000).UTC(), format: TimeFormatUnixDecimal}, expect: []byte(`0.999`)},
		{name: TimeFormatUnix, args: args{dst: make([]byte, 0), t: time.Unix(0, 999000000).UTC(), format: TimeFormatUnix}, expect: []byte(`0`)},
		{name: TimeFormatUnixMilli, args: args{dst: make([]byte, 0), t: time.Unix(0, 999000000).UTC(), format: TimeFormatUnixMilli}, expect: []byte(`999`)},
		{name: TimeFormatUnixMicro, args: args{dst: make([]byte, 0), t: time.Unix(0, 999000000).UTC(), format: TimeFormatUnixMicro}, expect: []byte(`999000`)},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := appendTimeFieldValue(tt.args.dst, tt.args.t, tt.args.format)
			FailIfNotBytesEqual(t, tt.expect, actual)
		})
	}
}
