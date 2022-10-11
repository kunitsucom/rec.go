// nolint: testpackage
package rec

import (
	"bytes"
	"io"
	"math"
	"math/big"
	"net/http"
	"testing"
	"time"
)

func TestBool(t *testing.T) {
	t.Parallel()

	t.Run("success(true)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":true`)
		actual := appendJSONField(bs, Bool("test", true))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(false)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":false`)
		actual := appendJSONField(bs, Bool("test", false))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestBoolPtr(t *testing.T) {
	t.Parallel()

	t.Run("success(true)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":true`)
		actual := appendJSONField(bs, BoolPtr("test", ptr.Bool(true)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(false)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":false`)
		actual := appendJSONField(bs, BoolPtr("test", ptr.Bool(false)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, BoolPtr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestUint(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxUint)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":18446744073709551615`)
		actual := appendJSONField(bs, Uint("test", math.MaxUint))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(0)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":0`)
		actual := appendJSONField(bs, Uint("test", 0))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestUintPtr(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxUint)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":18446744073709551615`)
		actual := appendJSONField(bs, UintPtr("test", ptr.Uint(math.MaxUint)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(0)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":0`)
		actual := appendJSONField(bs, UintPtr("test", ptr.Uint(0)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, UintPtr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestUint8(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxUint8)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":255`)
		actual := appendJSONField(bs, Uint8("test", math.MaxUint8))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(0)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":0`)
		actual := appendJSONField(bs, Uint8("test", 0))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestUint8Ptr(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxUint8)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":255`)
		actual := appendJSONField(bs, Uint8Ptr("test", ptr.Uint8(math.MaxUint8)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(0)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":0`)
		actual := appendJSONField(bs, Uint8Ptr("test", ptr.Uint8(0)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Uint8Ptr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestUint16(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxUint16)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":65535`)
		actual := appendJSONField(bs, Uint16("test", math.MaxUint16))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(0)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":0`)
		actual := appendJSONField(bs, Uint16("test", 0))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestUint16Ptr(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxUint16)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":65535`)
		actual := appendJSONField(bs, Uint16Ptr("test", ptr.Uint16(math.MaxUint16)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(0)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":0`)
		actual := appendJSONField(bs, Uint16Ptr("test", ptr.Uint16(0)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Uint16Ptr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestUint32(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxUint32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":4294967295`)
		actual := appendJSONField(bs, Uint32("test", math.MaxUint32))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(0)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":0`)
		actual := appendJSONField(bs, Uint32("test", 0))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestUint32Ptr(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxUint32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":4294967295`)
		actual := appendJSONField(bs, Uint32Ptr("test", ptr.Uint32(math.MaxUint32)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(0)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":0`)
		actual := appendJSONField(bs, Uint32Ptr("test", ptr.Uint32(0)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Uint32Ptr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestUint64(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxUint32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":18446744073709551615`)
		actual := appendJSONField(bs, Uint64("test", math.MaxUint64))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(0)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":0`)
		actual := appendJSONField(bs, Uint64("test", 0))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestUint64Ptr(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxUint64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":18446744073709551615`)
		actual := appendJSONField(bs, Uint64Ptr("test", ptr.Uint64(math.MaxUint64)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(0)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":0`)
		actual := appendJSONField(bs, Uint64Ptr("test", ptr.Uint64(0)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Uint64Ptr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestInt(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxInt)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":9223372036854775807`)
		actual := appendJSONField(bs, Int("test", math.MaxInt))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.MinInt)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":-9223372036854775808`)
		actual := appendJSONField(bs, Int("test", math.MinInt))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestIntPtr(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxInt)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":9223372036854775807`)
		actual := appendJSONField(bs, IntPtr("test", ptr.Int(math.MaxInt)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.MinInt)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":-9223372036854775808`)
		actual := appendJSONField(bs, IntPtr("test", ptr.Int(math.MinInt)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, IntPtr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestInt8(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxInt8)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":127`)
		actual := appendJSONField(bs, Int8("test", math.MaxInt8))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.MinInt8)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":-128`)
		actual := appendJSONField(bs, Int8("test", math.MinInt8))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestInt8Ptr(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxInt8)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":127`)
		actual := appendJSONField(bs, Int8Ptr("test", ptr.Int8(math.MaxInt8)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.MinInt8)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":-128`)
		actual := appendJSONField(bs, Int8Ptr("test", ptr.Int8(math.MinInt8)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Int8Ptr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestInt16(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxInt16)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":32767`)
		actual := appendJSONField(bs, Int16("test", math.MaxInt16))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.MinInt16)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":-32768`)
		actual := appendJSONField(bs, Int16("test", math.MinInt16))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestInt16Ptr(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxInt16)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":32767`)
		actual := appendJSONField(bs, Int16Ptr("test", ptr.Int16(math.MaxInt16)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.MinInt16)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":-32768`)
		actual := appendJSONField(bs, Int16Ptr("test", ptr.Int16(math.MinInt16)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Int16Ptr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestInt32(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxInt32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":2147483647`)
		actual := appendJSONField(bs, Int32("test", math.MaxInt32))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.MinInt32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":-2147483648`)
		actual := appendJSONField(bs, Int32("test", math.MinInt32))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestInt32Ptr(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxInt32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":2147483647`)
		actual := appendJSONField(bs, Int32Ptr("test", ptr.Int32(math.MaxInt32)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.MinInt32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":-2147483648`)
		actual := appendJSONField(bs, Int32Ptr("test", ptr.Int32(math.MinInt32)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Int32Ptr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestInt64(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxInt64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":9223372036854775807`)
		actual := appendJSONField(bs, Int64("test", math.MaxInt64))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.MinInt64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":-9223372036854775808`)
		actual := appendJSONField(bs, Int64("test", math.MinInt64))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestInt64Ptr(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxInt64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":9223372036854775807`)
		actual := appendJSONField(bs, Int64Ptr("test", ptr.Int64(math.MaxInt64)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.MinInt64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":-9223372036854775808`)
		actual := appendJSONField(bs, Int64Ptr("test", ptr.Int64(math.MinInt64)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Int64Ptr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestFloat32(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxFloat32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":340282350000000000000000000000000000000`)
		actual := appendJSONField(bs, Float32("test", math.MaxFloat32))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.SmallestNonzeroFloat32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":0.000000000000000000000000000000000000000000001`)
		actual := appendJSONField(bs, Float32("test", math.SmallestNonzeroFloat32))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestFloat32Ptr(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxFloat32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":340282350000000000000000000000000000000`)
		actual := appendJSONField(bs, Float32Ptr("test", ptr.Float32(math.MaxFloat32)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.SmallestNonzeroFloat32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":0.000000000000000000000000000000000000000000001`)
		actual := appendJSONField(bs, Float32Ptr("test", ptr.Float32(math.SmallestNonzeroFloat32)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Float32Ptr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestFloat64(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxFloat64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000`)
		actual := appendJSONField(bs, Float64("test", math.MaxFloat64))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.SmallestNonzeroFloat64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005`)
		actual := appendJSONField(bs, Float64("test", math.SmallestNonzeroFloat64))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(NaN)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"NaN"`)
		actual := appendJSONField(bs, Float64("test", math.NaN()))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(+Inf)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"+Inf"`)
		actual := appendJSONField(bs, Float64("test", math.Inf(1)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(-Inf)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"-Inf"`)
		actual := appendJSONField(bs, Float64("test", math.Inf(-1)))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestFloat64Ptr(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxFloat64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000`)
		actual := appendJSONField(bs, Float64Ptr("test", ptr.Float64(math.MaxFloat64)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.SmallestNonzeroFloat64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005`)
		actual := appendJSONField(bs, Float64Ptr("test", ptr.Float64(math.SmallestNonzeroFloat64)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(NaN)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"NaN"`)
		actual := appendJSONField(bs, Float64Ptr("test", ptr.Float64(math.NaN())))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(+Inf)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"+Inf"`)
		actual := appendJSONField(bs, Float64Ptr("test", ptr.Float64(math.Inf(1))))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(-Inf)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"-Inf"`)
		actual := appendJSONField(bs, Float64Ptr("test", ptr.Float64(math.Inf(-1))))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Float64Ptr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestComplex64(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxFloat32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"340282350000000000000000000000000000000+340282350000000000000000000000000000000i"`)
		actual := appendJSONField(bs, Complex64("test", complex(math.MaxFloat32, math.MaxFloat32)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(-math.MaxFloat32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"-340282350000000000000000000000000000000-340282350000000000000000000000000000000i"`)
		actual := appendJSONField(bs, Complex64("test", complex(-math.MaxFloat32, -math.MaxFloat32)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.SmallestNonzeroFloat32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"0.000000000000000000000000000000000000000000001+0.000000000000000000000000000000000000000000001i"`)
		actual := appendJSONField(bs, Complex64("test", complex(math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(-math.SmallestNonzeroFloat32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"-0.000000000000000000000000000000000000000000001-0.000000000000000000000000000000000000000000001i"`)
		actual := appendJSONField(bs, Complex64("test", complex(-math.SmallestNonzeroFloat32, -math.SmallestNonzeroFloat32)))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestComplex64Ptr(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxFloat32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"340282350000000000000000000000000000000+340282350000000000000000000000000000000i"`)
		actual := appendJSONField(bs, Complex64Ptr("test", ptr.Complex64(complex(math.MaxFloat32, math.MaxFloat32))))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(-math.MaxFloat32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"-340282350000000000000000000000000000000-340282350000000000000000000000000000000i"`)
		actual := appendJSONField(bs, Complex64Ptr("test", ptr.Complex64(complex(-math.MaxFloat32, -math.MaxFloat32))))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.SmallestNonzeroFloat32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"0.000000000000000000000000000000000000000000001+0.000000000000000000000000000000000000000000001i"`)
		actual := appendJSONField(bs, Complex64Ptr("test", ptr.Complex64(complex(math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32))))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(-math.SmallestNonzeroFloat32)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"-0.000000000000000000000000000000000000000000001-0.000000000000000000000000000000000000000000001i"`)
		actual := appendJSONField(bs, Complex64Ptr("test", ptr.Complex64(complex(-math.SmallestNonzeroFloat32, -math.SmallestNonzeroFloat32))))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Complex64Ptr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestComplex128(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxFloat64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000+179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000i"`)
		actual := appendJSONField(bs, Complex128("test", complex(math.MaxFloat64, math.MaxFloat64)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(-math.MaxFloat64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"-179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000-179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000i"`)
		actual := appendJSONField(bs, Complex128("test", complex(-math.MaxFloat64, -math.MaxFloat64)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.SmallestNonzeroFloat64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005+0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005i"`)
		actual := appendJSONField(bs, Complex128("test", complex(math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(-math.SmallestNonzeroFloat64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"-0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005-0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005i"`)
		actual := appendJSONField(bs, Complex128("test", complex(-math.SmallestNonzeroFloat64, -math.SmallestNonzeroFloat64)))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestComplex128Ptr(t *testing.T) {
	t.Parallel()

	t.Run("success(math.MaxFloat64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000+179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000i"`)
		actual := appendJSONField(bs, Complex128Ptr("test", ptr.Complex128(complex(math.MaxFloat64, math.MaxFloat64))))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(-math.MaxFloat64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"-179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000-179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000i"`)
		actual := appendJSONField(bs, Complex128Ptr("test", ptr.Complex128(complex(-math.MaxFloat64, -math.MaxFloat64))))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(math.SmallestNonzeroFloat64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005+0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005i"`)
		actual := appendJSONField(bs, Complex128Ptr("test", ptr.Complex128(complex(math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64))))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(-math.SmallestNonzeroFloat64)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"-0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005-0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005i"`)
		actual := appendJSONField(bs, Complex128Ptr("test", ptr.Complex128(complex(-math.SmallestNonzeroFloat64, -math.SmallestNonzeroFloat64))))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Complex128Ptr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

// nolint: paralleltest
func TestTime(t *testing.T) {
	t.Run("success(1970-01-01T09:00:00+09:00)", func(t *testing.T) {
		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"1970-01-01T09:00:00+09:00"`)
		actual := appendJSONField(bs, Time("test", time.Unix(0, 0).In(tzJST)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(0001-01-01T00:00:00Z)", func(t *testing.T) {
		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"0001-01-01T00:00:00Z"`)
		actual := appendJSONField(bs, Time("test", time.Time{}))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(0000)", func(t *testing.T) {
		bs := make([]byte, 0, 1024)

		backup := CustomTimeFormat
		t.Cleanup(func() { CustomTimeFormat = backup }) // nolint: paralleltest
		CustomTimeFormat = "1504"                       // nolint: paralleltest

		expect := []byte(`"test":"0000"`)
		actual := appendJSONField(bs, Time("test", time.Unix(0, 0).UTC()))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nilTimezoneField)", func(t *testing.T) {
		bs := make([]byte, 0, 1024)

		backup := CustomTimeFormat
		t.Cleanup(func() { CustomTimeFormat = backup }) // nolint: paralleltest
		CustomTimeFormat = "1504"                       // nolint: paralleltest

		expect := []byte(`"test":"0000"`)
		nilTimezoneField := Field{
			t:           typeTime,
			key:         "test",
			int64value1: time.Unix(0, 0).Unix(),
		}
		actual := appendJSONField(bs, nilTimezoneField)

		FailIfNotBytesEqual(t, expect, actual)
	})
}

// nolint: paralleltest
func TestTimePtr(t *testing.T) {
	t.Run("success(1970-01-01T09:00:00+09:00)", func(t *testing.T) {
		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"1970-01-01T09:00:00+09:00"`)
		actual := appendJSONField(bs, TimePtr("test", ptr.Time(time.Unix(0, 0).In(tzJST))))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(0001-01-01T00:00:00Z)", func(t *testing.T) {
		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"0001-01-01T00:00:00Z"`)
		actual := appendJSONField(bs, TimePtr("test", ptr.Time(time.Time{})))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(0000)", func(t *testing.T) {
		bs := make([]byte, 0, 1024)

		backup := CustomTimeFormat
		t.Cleanup(func() { CustomTimeFormat = backup }) // nolint: paralleltest
		CustomTimeFormat = "1504"                       // nolint: paralleltest

		expect := []byte(`"test":"0000"`)
		actual := appendJSONField(bs, TimePtr("test", ptr.Time(time.Unix(0, 0).UTC())))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, TimePtr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestTimeFormat(t *testing.T) {
	t.Parallel()

	t.Run("success(0000)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"0000"`)
		actual := appendJSONField(bs, TimeFormat("test", "1504", time.Unix(0, 0).UTC()))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nilTimezoneField)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"0000"`)
		nilTimezoneField := Field{
			t:            typeTimeFormat,
			key:          "test",
			int64value1:  time.Unix(0, 0).Unix(),
			stringvalue1: "1504",
		}
		actual := appendJSONField(bs, nilTimezoneField)

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestTimeFormatPtr(t *testing.T) {
	t.Parallel()

	t.Run("success(0000)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"0000"`)
		actual := appendJSONField(bs, TimeFormatPtr("test", "1504", ptr.Time(time.Unix(0, 0).UTC())))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, TimeFormatPtr("test", "1504", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestDuration(t *testing.T) {
	t.Parallel()

	t.Run("success(3661)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":3661`)
		actual := appendJSONField(bs, Duration("test", time.Second, time.Hour+time.Minute+time.Second+time.Millisecond+time.Microsecond+time.Nanosecond))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestDurationPtr(t *testing.T) {
	t.Parallel()

	t.Run("success(3661001001001)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":3661001001001`)
		actual := appendJSONField(bs, DurationPtr("test", time.Nanosecond, ptr.Duration(time.Hour+time.Minute+time.Second+time.Millisecond+time.Microsecond+time.Nanosecond)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, DurationPtr("test", time.Second, nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestDurationFormat(t *testing.T) {
	t.Parallel()

	t.Run("success(1h1m1.001001001s)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"1h1m1.001001001s"`)
		actual := appendJSONField(bs, DurationFormat("test", time.Hour+time.Minute+time.Second+time.Millisecond+time.Microsecond+time.Nanosecond))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestDurationFormatPtr(t *testing.T) {
	t.Parallel()

	t.Run("success(1h1m1.001001001s)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"1h1m1.001001001s"`)
		actual := appendJSONField(bs, DurationFormatPtr("test", ptr.Duration(time.Hour+time.Minute+time.Second+time.Millisecond+time.Microsecond+time.Nanosecond)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, DurationFormatPtr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestString(t *testing.T) {
	t.Parallel()

	t.Run("success(str)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"str"`)
		actual := appendJSONField(bs, String("test", "str"))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestSprintf(t *testing.T) {
	t.Parallel()

	t.Run("success(1970-01-01 00:00:00 +0000 UTC)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"EOF"`)
		actual := appendJSONField(bs, Sprintf("test", "%v", io.EOF))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"<nil>"`)
		actual := appendJSONField(bs, Sprintf("test", "%v", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestStringPtr(t *testing.T) {
	t.Parallel()

	t.Run("success(str)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"str"`)
		actual := appendJSONField(bs, StringPtr("test", ptr.String("str")))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, StringPtr("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestStrings(t *testing.T) {
	t.Parallel()

	t.Run("success(str)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":["str0","str1"]`)
		actual := appendJSONField(bs, Strings("test", []string{"str0", "str1"}))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Strings("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success([]string{})", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":[]`)
		actual := appendJSONField(bs, Strings("test", []string{}))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestStringer(t *testing.T) {
	t.Parallel()

	t.Run("success(1970-01-01 00:00:00 +0000 UTC)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"1970-01-01 00:00:00 +0000 UTC"`)
		actual := appendJSONField(bs, Stringer("test", time.Unix(0, 0).UTC()))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Stringer("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestFormatter(t *testing.T) {
	t.Parallel()

	t.Run("success(+9223372036854775807)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"+9223372036854775807"`)
		actual := appendJSONField(bs, Formatter("test", big.NewInt(math.MaxInt64)))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Formatter("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestError(t *testing.T) {
	t.Parallel()

	t.Run("success(error)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"error":"EOF"`)
		actual := appendJSONField(bs, Error(io.EOF))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"error":null`)
		actual := appendJSONField(bs, Error(nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestErrorWithKey(t *testing.T) {
	t.Parallel()

	t.Run("success(error)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"EOF"`)
		actual := appendJSONField(bs, ErrorWithKey("test", io.EOF))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, ErrorWithKey("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestErrors(t *testing.T) {
	t.Parallel()

	t.Run("success(errors)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"errors":["EOF","unexpected EOF"]`)
		actual := appendJSONField(bs, Errors([]error{io.EOF, io.ErrUnexpectedEOF}))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"errors":null`)
		actual := appendJSONField(bs, Errors(nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestErrorsWithKey(t *testing.T) {
	t.Parallel()

	t.Run("success(errors)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":["EOF","http: connection has been hijacked"]`)
		actual := appendJSONField(bs, ErrorsWithKey("test", []error{io.EOF, http.ErrHijacked}))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, ErrorsWithKey("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success([]error{})", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":[]`)
		actual := appendJSONField(bs, ErrorsWithKey("test", []error{}))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

type testFormaterError struct {
	err error
	big.Float
}

func (e testFormaterError) Error() string {
	return e.err.Error()
}

var testFormatErrorValue = &testFormaterError{err: io.EOF, Float: *big.NewFloat(math.Inf(-1))} // nolint: errname

func TestErrorStacktrace(t *testing.T) {
	t.Parallel()

	t.Run("success(errorStacktrace)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"errorStacktrace":"-Inf"`)
		actual := appendJSONField(bs, ErrorStacktrace(testFormatErrorValue))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(error)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"errorStacktrace":"EOF"`)
		actual := appendJSONField(bs, ErrorStacktrace(io.EOF))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"errorStacktrace":null`)
		actual := appendJSONField(bs, ErrorStacktrace(nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

func TestInterface(t *testing.T) {
	t.Parallel()

	t.Run("success(interface{})", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":"EOF"`)
		actual := appendJSONField(bs, Interface("test", io.EOF))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		t.Parallel()

		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Interface("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})
}

type testJSONObject struct {
	Test  bool    `json:"test"`
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

// nolint: paralleltest
func TestObject(t *testing.T) {
	t.Run("success(object)", func(t *testing.T) {
		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":{"test":true,"key":"test","value":3.141592653589793}`)
		actual := appendJSONField(bs, Object("test", testJSONObject{Test: true, Key: "test", Value: math.Pi}))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("success(nil)", func(t *testing.T) {
		bs := make([]byte, 0, 1024)

		expect := []byte(`"test":null`)
		actual := appendJSONField(bs, Object("test", nil))

		FailIfNotBytesEqual(t, expect, actual)
	})

	t.Run("error(unsupported)", func(t *testing.T) {
		backup := defaultLogger

		t.Cleanup(func() { defaultLogger = backup }) // nolint: paralleltest

		buf := bytes.NewBuffer(nil)
		defaultLogger = Must(New(buf, WithUseTimestampField(false), WithUseCallerField(false))) // nolint: paralleltest

		defaultLogger.Info("test", Object("unsupported", http.Request{Method: http.MethodGet}))

		expect := `{"severity":"ERROR","message":"rec.Object: json.Marshal: json: unsupported type: func() (io.ReadCloser, error)","error":"json: unsupported type: func() (io.ReadCloser, error)"}` + defaultLineSeparator + `{"severity":"INFO","message":"test","unsupported":null}` + defaultLineSeparator
		actual := buf.String()

		FailIfNotEqual(t, expect, actual)
	})
}
