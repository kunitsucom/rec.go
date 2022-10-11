package rec

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"
)

const (
	errorKey           = "error"
	errorsKey          = "errors"
	errorStacktraceKey = "errorStacktrace"
)

// Type is enum for rec.Field.
type Type uint8

const (
	typeNone Type = iota
	typeBool
	typeBoolPtr
	typeUint
	typeUintPtr
	typeUint8
	typeUint8Ptr
	typeUint16
	typeUint16Ptr
	typeUint32
	typeUint32Ptr
	typeUint64
	typeUint64Ptr
	typeInt
	typeIntPtr
	typeInt8
	typeInt8Ptr
	typeInt16
	typeInt16Ptr
	typeInt32
	typeInt32Ptr
	typeInt64
	typeInt64Ptr
	typeFloat32
	typeFloat32Ptr
	typeFloat64
	typeFloat64Ptr
	typeComplex64
	typeComplex64Ptr
	typeComplex128
	typeComplex128Ptr
	typeTime
	typeTimeFormat
	typeTimePtr
	typeTimeFormatPtr
	typeDuration
	typeDurationPtr
	typeDurationFormat
	typeDurationFormatPtr
	typeString
	typeStrings
	typeStringPtr
	typeStringer
	typeFormatter
	typeError
	typeErrors
	typeErrorStacktrace
	typeInterface
	typeObject
)

// DefaultTimeFormat is default time format for rec.Time() and rec.TimePtr().
const DefaultTimeFormat = time.RFC3339Nano

// CustomTimeFormat is time format for rec.Time() and rec.TimePtr().
var CustomTimeFormat = DefaultTimeFormat // nolint: gochecknoglobals

// Field is a struct for adding fields to the rec JSON log.
type Field struct {
	t               Type
	key             string
	interfacevalue1 interface{}
	boolvalue1      bool
	int64value1     int64
	int64value2     int64
	uint64value1    uint64
	float64value1   float64
	float64value2   float64
	stringvalue1    string
}

func appendJSONField(dst []byte, f Field) []byte {
	dst = append(dst, '"')
	dst = appendJSONEscapedString(dst, f.key)
	dst = append(dst, '"')
	dst = append(dst, ':')
	dst = appendFieldValue(dst, f, json.Marshal)

	return dst
}

// nolint: cyclop, funlen, gocognit, gocyclo, maintidx
func appendFieldValue(dst []byte, f Field, jsonMarshalFn func(interface{}) ([]byte, error)) []byte {
	const (
		b10  = 10
		b32  = 32
		b64  = 64
		null = `null`
	)

	switch f.t {
	// bool
	case typeBool:
		dst = strconv.AppendBool(dst, f.boolvalue1)
	case typeBoolPtr:
		value, ok := f.interfacevalue1.(*bool)
		if ok && value != nil {
			dst = strconv.AppendBool(dst, *value)

			break
		}

		dst = append(dst, null...)
	// number
	case typeUint:
		dst = strconv.AppendUint(dst, f.uint64value1, b10)
	case typeUintPtr:
		value, ok := f.interfacevalue1.(*uint)
		if ok && value != nil {
			dst = strconv.AppendUint(dst, uint64(*value), b10)

			break
		}

		dst = append(dst, null...)
	case typeUint8:
		dst = strconv.AppendUint(dst, f.uint64value1, b10)
	case typeUint8Ptr:
		value, ok := f.interfacevalue1.(*uint8)
		if ok && value != nil {
			dst = strconv.AppendUint(dst, uint64(*value), b10)

			break
		}

		dst = append(dst, null...)
	case typeUint16:
		dst = strconv.AppendUint(dst, f.uint64value1, b10)
	case typeUint16Ptr:
		value, ok := f.interfacevalue1.(*uint16)
		if ok && value != nil {
			dst = strconv.AppendUint(dst, uint64(*value), b10)

			break
		}

		dst = append(dst, null...)
	case typeUint32:
		dst = strconv.AppendUint(dst, f.uint64value1, b10)
	case typeUint32Ptr:
		value, ok := f.interfacevalue1.(*uint32)
		if ok && value != nil {
			dst = strconv.AppendUint(dst, uint64(*value), b10)

			break
		}

		dst = append(dst, null...)
	case typeUint64:
		dst = strconv.AppendUint(dst, f.uint64value1, b10)
	case typeUint64Ptr:
		value, ok := f.interfacevalue1.(*uint64)
		if ok && value != nil {
			dst = strconv.AppendUint(dst, *value, b10)

			break
		}

		dst = append(dst, null...)
	case typeInt:
		dst = strconv.AppendInt(dst, f.int64value1, b10)
	case typeIntPtr:
		value, ok := f.interfacevalue1.(*int)
		if ok && value != nil {
			dst = strconv.AppendInt(dst, int64(*value), b10)

			break
		}

		dst = append(dst, null...)
	case typeInt8:
		dst = strconv.AppendInt(dst, f.int64value1, b10)
	case typeInt8Ptr:
		value, ok := f.interfacevalue1.(*int8)
		if ok && value != nil {
			dst = strconv.AppendInt(dst, int64(*value), b10)

			break
		}

		dst = append(dst, null...)
	case typeInt16:
		dst = strconv.AppendInt(dst, f.int64value1, b10)
	case typeInt16Ptr:
		value, ok := f.interfacevalue1.(*int16)
		if ok && value != nil {
			dst = strconv.AppendInt(dst, int64(*value), b10)

			break
		}

		dst = append(dst, null...)
	case typeInt32:
		dst = strconv.AppendInt(dst, f.int64value1, b10)
	case typeInt32Ptr:
		value, ok := f.interfacevalue1.(*int32)
		if ok && value != nil {
			dst = strconv.AppendInt(dst, int64(*value), b10)

			break
		}

		dst = append(dst, null...)
	case typeInt64:
		dst = strconv.AppendInt(dst, f.int64value1, b10)
	case typeInt64Ptr:
		value, ok := f.interfacevalue1.(*int64)
		if ok && value != nil {
			dst = strconv.AppendInt(dst, *value, b10)

			break
		}

		dst = append(dst, null...)
	case typeFloat32:
		dst = appendFloatFieldValue(dst, f.float64value1, b32)
	case typeFloat32Ptr:
		value, ok := f.interfacevalue1.(*float32)
		if ok && value != nil {
			dst = appendFloatFieldValue(dst, float64(*value), b32)

			break
		}

		dst = append(dst, null...)
	case typeFloat64:
		dst = appendFloatFieldValue(dst, f.float64value1, b64)
	case typeFloat64Ptr:
		value, ok := f.interfacevalue1.(*float64)
		if ok && value != nil {
			dst = appendFloatFieldValue(dst, *value, b64)

			break
		}

		dst = append(dst, null...)
	// string
	case typeComplex64:
		dst = append(dst, '"')
		dst = strconv.AppendFloat(dst, f.float64value1, 'f', -1, b32)

		if f.float64value2 >= 0 && !math.IsInf(f.float64value2, b32) {
			dst = append(dst, '+')
		}

		dst = strconv.AppendFloat(dst, f.float64value2, 'f', -1, b32)
		dst = append(dst, 'i')
		dst = append(dst, '"')
	case typeComplex64Ptr:
		value, ok := f.interfacevalue1.(*complex64)

		if ok && value != nil {
			dst = append(dst, '"')
			dst = strconv.AppendFloat(dst, float64(real(*value)), 'f', -1, b32)
			imaginary := imag(*value)

			if imaginary >= 0 && !math.IsInf(float64(imaginary), b32) {
				dst = append(dst, '+')
			}

			dst = strconv.AppendFloat(dst, float64(imaginary), 'f', -1, b32)
			dst = append(dst, 'i')
			dst = append(dst, '"')

			break
		}

		dst = append(dst, null...)
	case typeComplex128:
		dst = append(dst, '"')
		dst = strconv.AppendFloat(dst, f.float64value1, 'f', -1, b64)

		if f.float64value2 >= 0 && !math.IsInf(f.float64value2, b64) {
			dst = append(dst, '+')
		}

		dst = strconv.AppendFloat(dst, f.float64value2, 'f', -1, b64)
		dst = append(dst, 'i')
		dst = append(dst, '"')
	case typeComplex128Ptr:
		value, ok := f.interfacevalue1.(*complex128)

		if ok && value != nil {
			dst = append(dst, '"')
			dst = strconv.AppendFloat(dst, real(*value), 'f', -1, b64)
			imaginary := imag(*value)

			if imaginary >= 0 && !math.IsInf(imaginary, b64) {
				dst = append(dst, '+')
			}

			dst = strconv.AppendFloat(dst, imaginary, 'f', -1, b64)
			dst = append(dst, 'i')
			dst = append(dst, '"')

			break
		}

		dst = append(dst, null...)
	case typeTime:
		value, _ := f.interfacevalue1.(*time.Location)

		if value == nil {
			value = time.UTC
		}

		dst = appendTimeFieldValue(dst, time.Unix(f.int64value1, f.int64value2).In(value), CustomTimeFormat)
	case typeTimeFormat:
		value, _ := f.interfacevalue1.(*time.Location)

		if value == nil {
			value = time.UTC
		}

		dst = appendTimeFieldValue(dst, time.Unix(f.int64value1, f.int64value2).In(value), f.stringvalue1)
	case typeTimePtr:
		value, ok := f.interfacevalue1.(*time.Time)

		if ok && value != nil {
			dst = appendTimeFieldValue(dst, *value, CustomTimeFormat)

			break
		}

		dst = append(dst, null...)
	case typeTimeFormatPtr:
		value, ok := f.interfacevalue1.(*time.Time)

		if ok && value != nil {
			dst = appendTimeFieldValue(dst, *value, f.stringvalue1)

			break
		}

		dst = append(dst, null...)
	case typeDuration:
		dst = strconv.AppendInt(dst, f.int64value2/f.int64value1, b10)
	case typeDurationPtr:
		value, ok := f.interfacevalue1.(*time.Duration)

		if ok && value != nil {
			dst = strconv.AppendInt(dst, int64(*value)/f.int64value1, b10)

			break
		}

		dst = append(dst, null...)
	case typeDurationFormat:
		dst = append(appendJSONEscapedString(append(dst, '"'), f.stringvalue1), '"')
	case typeDurationFormatPtr:
		value, ok := f.interfacevalue1.(*time.Duration)

		if ok && value != nil {
			dst = append(appendJSONEscapedString(append(dst, '"'), value.String()), '"')

			break
		}

		dst = append(dst, null...)
	case typeString:
		dst = append(appendJSONEscapedString(append(dst, '"'), f.stringvalue1), '"')
	case typeStringPtr:
		value, ok := f.interfacevalue1.(*string)

		if ok && value != nil {
			dst = append(appendJSONEscapedString(append(dst, '"'), *value), '"')

			break
		}

		dst = append(dst, null...)
	case typeStrings:
		value, ok := f.interfacevalue1.([]string)

		if ok && value != nil {
			dst = append(dst, '[')

			for _, v := range value {
				dst = append(append(appendJSONEscapedString(append(dst, '"'), v), '"'), ',')
			}

			if dst[len(dst)-1] == ',' {
				dst[len(dst)-1] = ']'
			}

			if dst[len(dst)-1] != ']' {
				dst = append(dst, ']')
			}

			break
		}

		dst = append(dst, null...)
	case typeStringer:
		value, ok := f.interfacevalue1.(fmt.Stringer)

		if ok && value != nil {
			dst = append(appendJSONEscapedString(append(dst, '"'), value.String()), '"')

			break
		}

		dst = append(dst, null...)
	case typeFormatter:
		value, ok := f.interfacevalue1.(fmt.Formatter)
		if ok && value != nil {
			dst = append(appendJSONEscapedString(append(dst, '"'), fmt.Sprintf("%+v", value)), '"')

			break
		}

		dst = append(dst, null...)
	case typeError:
		value, ok := f.interfacevalue1.(error)

		if ok && value != nil {
			dst = append(appendJSONEscapedString(append(dst, '"'), value.Error()), '"')

			break
		}

		dst = append(dst, null...)
	case typeErrors:
		value, ok := f.interfacevalue1.([]error)

		if ok && value != nil {
			dst = append(dst, '[')

			for _, v := range value {
				dst = append(append(appendJSONEscapedString(append(dst, '"'), v.Error()), '"'), ',')
			}

			if dst[len(dst)-1] == ',' {
				dst[len(dst)-1] = ']'
			}

			if dst[len(dst)-1] != ']' {
				dst = append(dst, ']')
			}

			break
		}

		dst = append(dst, null...)
	case typeErrorStacktrace:
		value, ok := f.interfacevalue1.(fmt.Formatter)
		if ok && value != nil {
			dst = append(appendJSONEscapedString(append(dst, '"'), fmt.Sprintf("%+v", value)), '"')

			break
		}

		err, ok := f.interfacevalue1.(error)
		if ok && err != nil {
			dst = append(appendJSONEscapedString(append(dst, '"'), err.Error()), '"')

			break
		}

		dst = append(dst, null...)
	case typeInterface:
		value := fmt.Sprintf("%+v", f.interfacevalue1)

		if value != "<nil>" {
			dst = append(appendJSONEscapedString(append(dst, '"'), value), '"')

			break
		}

		dst = append(dst, null...)
	case typeObject:
		b, err := jsonMarshalFn(f.interfacevalue1)
		if err != nil {
			const skip = 4

			defaultLogger.AddCallerSkip(skip).Error("rec.Object: json.Marshal: "+err.Error(), Error(err))

			dst = append(dst, null...)

			break
		}

		dst = append(dst, b...)
	// abnormal
	case typeNone:
		dst = append(dst, `"ERROR: TYPE NONE"`...)
	default:
		dst = append(strconv.AppendInt(append(dst, `"ERROR: UNDEFINED TYPE: `...), int64(f.t), b10), '"')
	}

	return dst
}

// Bool returns rec.Field for bool type.
func Bool(key string, value bool) Field {
	return Field{
		t:          typeBool,
		key:        key,
		boolvalue1: value,
	}
}

// BoolPtr returns rec.Field for *bool type.
func BoolPtr(key string, value *bool) Field {
	return Field{
		t:               typeBoolPtr,
		key:             key,
		interfacevalue1: value,
	}
}

// Uint returns rec.Field for uint type.
func Uint(key string, value uint) Field {
	return Field{
		t:            typeUint,
		key:          key,
		uint64value1: uint64(value),
	}
}

// UintPtr returns rec.Field for *uint type.
func UintPtr(key string, value *uint) Field {
	return Field{
		t:               typeUintPtr,
		key:             key,
		interfacevalue1: value,
	}
}

// Uint8 returns rec.Field for uint8 type.
func Uint8(key string, value uint8) Field {
	return Field{
		t:            typeUint8,
		key:          key,
		uint64value1: uint64(value),
	}
}

// Uint8Ptr returns rec.Field for *uint8 type.
func Uint8Ptr(key string, value *uint8) Field {
	return Field{
		t:               typeUint8Ptr,
		key:             key,
		interfacevalue1: value,
	}
}

// Uint16 returns rec.Field for uint16 type.
func Uint16(key string, value uint16) Field {
	return Field{
		t:            typeUint16,
		key:          key,
		uint64value1: uint64(value),
	}
}

// Uint16Ptr returns rec.Field for *uint16 type.
func Uint16Ptr(key string, value *uint16) Field {
	return Field{
		t:               typeUint16Ptr,
		key:             key,
		interfacevalue1: value,
	}
}

// Uint32 returns rec.Field for uint32 type.
func Uint32(key string, value uint32) Field {
	return Field{
		t:            typeUint32,
		key:          key,
		uint64value1: uint64(value),
	}
}

// Uint32Ptr returns rec.Field for *uint32 type.
func Uint32Ptr(key string, value *uint32) Field {
	return Field{
		t:               typeUint32Ptr,
		key:             key,
		interfacevalue1: value,
	}
}

// Uint64 returns rec.Field for uint64 type.
func Uint64(key string, value uint64) Field {
	return Field{
		t:            typeUint64,
		key:          key,
		uint64value1: value,
	}
}

// Uint64Ptr returns rec.Field for *uint64 type.
func Uint64Ptr(key string, value *uint64) Field {
	return Field{
		t:               typeUint64Ptr,
		key:             key,
		interfacevalue1: value,
	}
}

// Int returns rec.Field for int type.
func Int(key string, value int) Field {
	return Field{
		t:           typeInt,
		key:         key,
		int64value1: int64(value),
	}
}

// IntPtr returns rec.Field for *int type.
func IntPtr(key string, value *int) Field {
	return Field{
		t:               typeIntPtr,
		key:             key,
		interfacevalue1: value,
	}
}

// Int8 returns rec.Field for int8 type.
func Int8(key string, value int8) Field {
	return Field{
		t:           typeInt8,
		key:         key,
		int64value1: int64(value),
	}
}

// Int8Ptr returns rec.Field for *int8 type.
func Int8Ptr(key string, value *int8) Field {
	return Field{
		t:               typeInt8Ptr,
		key:             key,
		interfacevalue1: value,
	}
}

// Int16 returns rec.Field for int16 type.
func Int16(key string, value int16) Field {
	return Field{
		t:           typeInt16,
		key:         key,
		int64value1: int64(value),
	}
}

// Int16Ptr returns rec.Field for *int16 type.
func Int16Ptr(key string, value *int16) Field {
	return Field{
		t:               typeInt16Ptr,
		key:             key,
		interfacevalue1: value,
	}
}

// Int32 returns rec.Field for int32 type.
func Int32(key string, value int32) Field {
	return Field{
		t:           typeInt32,
		key:         key,
		int64value1: int64(value),
	}
}

// Int32Ptr returns rec.Field for *int32 type.
func Int32Ptr(key string, value *int32) Field {
	return Field{
		t:               typeInt32Ptr,
		key:             key,
		interfacevalue1: value,
	}
}

// Int64 returns rec.Field for int64 type.
func Int64(key string, value int64) Field {
	return Field{
		t:           typeInt64,
		key:         key,
		int64value1: value,
	}
}

// Int64Ptr returns rec.Field for *int64 type.
func Int64Ptr(key string, value *int64) Field {
	return Field{
		t:               typeInt64Ptr,
		key:             key,
		interfacevalue1: value,
	}
}

// Float32 returns rec.Field for float32 type.
func Float32(key string, value float32) Field {
	return Field{
		t:             typeFloat32,
		key:           key,
		float64value1: float64(value),
	}
}

// Float32Ptr returns rec.Field for *float32 type.
func Float32Ptr(key string, value *float32) Field {
	return Field{
		t:               typeFloat32Ptr,
		key:             key,
		interfacevalue1: value,
	}
}

// Float64 returns rec.Field for float64 type.
func Float64(key string, value float64) Field {
	return Field{
		t:             typeFloat64,
		key:           key,
		float64value1: value,
	}
}

// Float64Ptr returns rec.Field for *float64 type.
func Float64Ptr(key string, value *float64) Field {
	return Field{
		t:               typeFloat64Ptr,
		key:             key,
		interfacevalue1: value,
	}
}

// Complex64 returns rec.Field for complex64 type.
func Complex64(key string, value complex64) Field {
	return Field{
		t:             typeComplex64,
		key:           key,
		float64value1: float64(real(value)),
		float64value2: float64(imag(value)),
	}
}

// Complex64Ptr returns rec.Field for *complex128 type.
func Complex64Ptr(key string, value *complex64) Field {
	return Field{
		t:               typeComplex64Ptr,
		key:             key,
		interfacevalue1: value,
	}
}

// Complex128 returns rec.Field for complex128 type.
func Complex128(key string, value complex128) Field {
	return Field{
		t:             typeComplex128,
		key:           key,
		float64value1: real(value),
		float64value2: imag(value),
	}
}

// Complex128Ptr returns rec.Field for *complex128 type.
func Complex128Ptr(key string, value *complex128) Field {
	return Field{
		t:               typeComplex128Ptr,
		key:             key,
		interfacevalue1: value,
	}
}

// Time returns rec.Field for time.Time type.
func Time(key string, value time.Time) Field {
	return Field{
		t:               typeTime,
		key:             key,
		interfacevalue1: value.Location(),
		int64value1:     value.Unix(),
		int64value2:     int64(value.Nanosecond()),
	}
}

// TimeFormat returns rec.Field for time.Time type with time format.
func TimeFormat(key string, format string, value time.Time) Field {
	return Field{
		t:               typeTimeFormat,
		key:             key,
		interfacevalue1: value.Location(),
		int64value1:     value.Unix(),
		int64value2:     int64(value.Nanosecond()),
		stringvalue1:    format,
	}
}

// TimePtr returns rec.Field for *time.Time type.
func TimePtr(key string, value *time.Time) Field {
	return Field{
		t:               typeTimePtr,
		key:             key,
		interfacevalue1: value,
	}
}

// TimeFormatPtr returns rec.Field for *time.Time type with time format.
func TimeFormatPtr(key string, format string, value *time.Time) Field {
	return Field{
		t:               typeTimeFormatPtr,
		key:             key,
		interfacevalue1: value,
		stringvalue1:    format,
	}
}

// Duration returns rec.Field for time.Duration type.
func Duration(key string, unit time.Duration, value time.Duration) Field {
	return Field{
		t:           typeDuration,
		key:         key,
		int64value1: int64(unit),
		int64value2: int64(value),
	}
}

// DurationPtr returns rec.Field for time.Duration type.
func DurationPtr(key string, unit time.Duration, value *time.Duration) Field {
	return Field{
		t:               typeDurationPtr,
		key:             key,
		int64value1:     int64(unit),
		interfacevalue1: value,
	}
}

// DurationFormat returns rec.Field for time.Duration type with duration format.
func DurationFormat(key string, value time.Duration) Field {
	return Field{
		t:            typeDurationFormat,
		key:          key,
		stringvalue1: value.String(),
	}
}

// DurationFormatPtr returns rec.Field for time.Duration type with duration format.
func DurationFormatPtr(key string, value *time.Duration) Field {
	return Field{
		t:               typeDurationFormatPtr,
		key:             key,
		interfacevalue1: value,
	}
}

// String returns rec.Field for string type.
func String(key string, value string) Field {
	return Field{
		t:            typeString,
		key:          key,
		stringvalue1: value,
	}
}

// Sprintf returns rec.Field like fmt.Sprintf() for string type.
func Sprintf(key string, format string, a ...interface{}) Field {
	return Field{
		t:            typeString,
		key:          key,
		stringvalue1: fmt.Sprintf(format, a...),
	}
}

// StringPtr returns rec.Field like fmt.Sprintf() for *string type.
func StringPtr(key string, value *string) Field {
	return Field{
		t:               typeStringPtr,
		key:             key,
		interfacevalue1: value,
	}
}

// Strings returns rec.Field for []string type.
func Strings(key string, value []string) Field {
	return Field{
		t:               typeStrings,
		key:             key,
		interfacevalue1: value,
	}
}

// Stringer returns rec.Field for fmt.Stringer type.
func Stringer(key string, value fmt.Stringer) Field {
	return Field{
		t:               typeStringer,
		key:             key,
		interfacevalue1: value,
	}
}

// Formatter returns rec.Field for fmt.Formatter type.
func Formatter(key string, value fmt.Formatter) Field {
	return Field{
		t:               typeFormatter,
		key:             key,
		interfacevalue1: value,
	}
}

// Error returns rec.Field for error type.
// Field key is fixed to `error`.
// If you want to use other field key, use rec.ErrorWithKey().
func Error(err error) Field {
	return ErrorWithKey(errorKey, err)
}

// ErrorWithKey returns rec.Field for error type.
func ErrorWithKey(key string, err error) Field {
	return Field{
		t:               typeError,
		key:             key,
		interfacevalue1: err,
	}
}

// Errors returns rec.Field for []error type.
// Field key is fixed to `errors`.
// If you want to use other field key, use rec.ErrorsWithKey().
func Errors(errs []error) Field {
	return ErrorsWithKey(errorsKey, errs)
}

// ErrorsWithKey returns rec.Field for []error type.
func ErrorsWithKey(key string, errs []error) Field {
	return Field{
		t:               typeErrors,
		key:             key,
		interfacevalue1: errs,
	}
}

// ErrorStacktrace returns rec.Field for error type with stacktrace.
// Field key is fixed to `errorStacktrace`.
// If you want to use other field key, use rec.ErrorStacktraceWithKey().
func ErrorStacktrace(err error) Field {
	return ErrorStacktraceWithKey(errorStacktraceKey, err)
}

// ErrorStacktraceWithKey returns rec.Field for error type with stacktrace.
func ErrorStacktraceWithKey(key string, err error) Field {
	return Field{
		t:               typeErrorStacktrace,
		key:             key,
		interfacevalue1: err,
	}
}

// Interface returns rec.Field for interface{} type.
func Interface(key string, value interface{}) Field {
	return Field{
		t:               typeInterface,
		key:             key,
		interfacevalue1: value,
	}
}

// Object returns a rec.Field for the JSON Object struct as shown below:
//
//	type Example {
//	    Key   string `json:"key"`
//	    Value string `json:"value"`
//	}
func Object(key string, object interface{}) Field {
	return Field{
		t:               typeObject,
		key:             key,
		interfacevalue1: object,
	}
}
