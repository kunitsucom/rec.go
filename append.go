package rec

import (
	"math"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// apppendUnicodeEscapeSequence converts a single byte into bytes sequence of Unicode Escape Sequence and appends to dst.
func apppendUnicodeEscapeSequence(dst []byte, singleByte byte) []byte {
	const hextable string = "0123456789abcdef"
	// cf. https://github.com/golang/go/blob/70deaa33ebd91944484526ab368fa19c499ff29f/src/encoding/hex/hex.go#L28-L29
	dst = append(dst, '\\', 'u', '0', '0', hextable[singleByte>>4], hextable[singleByte&0x0f])

	return dst
}

// nolint: cyclop
// appendJSONEscapedString
func appendJSONEscapedString(dst []byte, s string) []byte {
	for i := 0; i < len(s); i++ {
		if s[i] != '"' && s[i] != '\\' && s[i] > 0x1F {
			dst = append(dst, s[i])

			continue
		}

		// cf. https://tools.ietf.org/html/rfc8259#section-7
		// ... MUST be escaped: quotation mark, reverse solidus, and the control characters (U+0000 through U+001F).
		switch s[i] {
		case '"', '\\':
			dst = append(dst, '\\', s[i])
		case '\b' /* 0x08 */ :
			dst = append(dst, '\\', 'b')
		case '\f' /* 0x0C */ :
			dst = append(dst, '\\', 'f')
		case '\n' /* 0x0A */ :
			dst = append(dst, '\\', 'n')
		case '\r' /* 0x0D */ :
			dst = append(dst, '\\', 'r')
		case '\t' /* 0x09 */ :
			dst = append(dst, '\\', 't')
		default:
			dst = apppendUnicodeEscapeSequence(dst, s[i])
		}
	}

	return dst
}

func appendFloatFieldValue(dst []byte, value float64, bitSize int) []byte {
	switch {
	case math.IsNaN(value):
		return append(dst, `"NaN"`...)
	case math.IsInf(value, 1):
		return append(dst, `"+Inf"`...)
	case math.IsInf(value, -1):
		return append(dst, `"-Inf"`...)
	}

	return strconv.AppendFloat(dst, value, 'f', -1, bitSize)
}

const (
	// TimeFormatUnixDecimal is time format for rec.Config.TimestampFieldFormat and rec.TimeFormat() and rec.TimeFormatPtr()
	// UNIXDECIMAL format style ... 1638645867.123456789
	TimeFormatUnixDecimal = "UNIXDECIMAL"
	// TimeFormatUnix is time format for rec.Config.TimestampFieldFormat and rec.TimeFormat() and rec.TimeFormatPtr()
	// UNIX format style ... 1638645867
	TimeFormatUnix = "UNIX"
	// TimeFormatUnixMilli is time format for rec.Config.TimestampFieldFormat and rec.TimeFormat() and rec.TimeFormatPtr()
	// UNIXMILLI format style ... 1638645867123
	TimeFormatUnixMilli = "UNIXMILLI"
	// TimeFormatUnixMicro is time format for rec.Config.TimestampFieldFormat and rec.TimeFormat() and rec.TimeFormatPtr()
	// UNIXMICRO format style ... 1638645867123456
	TimeFormatUnixMicro = "UNIXMICRO"
)

func appendTimeFieldValue(dst []byte, t time.Time, format string) []byte {
	const base = 10

	switch format {
	case "", TimeFormatUnixDecimal:
		dst = strconv.AppendInt(dst, t.Unix(), base)
		if nsec := t.Nanosecond(); nsec != 0 {
			dst = append(dst, '.')

			for i := base; nsec%i == 0; { // 0.100000000 -> 0.1
				nsec /= i
			}

			dst = strconv.AppendInt(dst, int64(nsec), base)
		}
	case TimeFormatUnix:
		dst = strconv.AppendInt(dst, t.Unix(), base)
	case TimeFormatUnixMilli:
		dst = strconv.AppendInt(dst, t.UnixMilli(), base)
	case TimeFormatUnixMicro:
		dst = strconv.AppendInt(dst, t.UnixMicro(), base)
	default:
		dst = append(t.AppendFormat(append(dst, '"'), format), '"')
	}

	return dst
}

func appendCaller(dst []byte, callerSkip int, useShortCaller bool) []byte {
	pc := pcPool.Get().(*programcounter) // nolint: forcetypeassert
	defer pcPool.Put(pc)

	var frame runtime.Frame
	if runtime.Callers(callerSkip, pc.PC) > 0 {
		frame, _ = runtime.CallersFrames(pc.PC).Next()
	}

	return appendCallerFromFrame(dst, frame, useShortCaller)
}

// appendCallerFromFrame was split off from appendCaller in order to test different behaviors depending on the contents of the `runtime.Frame`.
func appendCallerFromFrame(dst []byte, frame runtime.Frame, useShortCaller bool) []byte {
	const base = 10

	if useShortCaller {
		dst = appendJSONEscapedString(dst, extractShortPath(frame.File))
	} else {
		dst = appendJSONEscapedString(dst, frame.File)
	}

	dst = append(dst, ':')
	dst = strconv.AppendInt(dst, int64(frame.Line), base)

	return dst
}

func extractShortPath(path string) string {
	// path == /path/to/directory/file
	//                           ~ <- idx
	idx := strings.LastIndexByte(path, '/')
	if idx == -1 {
		return path
	}

	// path[:idx] == /path/to/directory
	//                       ~ <- idx
	idx = strings.LastIndexByte(path[:idx], '/')
	if idx == -1 {
		return path
	}

	// path == /path/to/directory/file
	//                  ~~~~~~~~~~~~~~ <- filepath[idx+1:]
	return path[idx+1:]
}
