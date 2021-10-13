// nolint: testpackage
package rec

import (
	"bytes"
	"reflect"
	"regexp"
	"testing"
)

func TestLogger_Print(t *testing.T) {
	t.Parallel()

	t.Run("success(Print)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.Print(DEFAULT, "test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"DEFAULT","host":"test","caller":"[^"]+:[0-9]+","message":"test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Default)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.Default("test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"DEFAULT","host":"test","caller":"[^"]+:[0-9]+","message":"test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Debug)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.Debug("test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"DEBUG","host":"test","caller":"[^"]+:[0-9]+","message":"test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Info)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.Info("test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"INFO","host":"test","caller":"[^"]+:[0-9]+","message":"test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Notice)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.Notice("test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"NOTICE","host":"test","caller":"[^"]+:[0-9]+","message":"test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Warning)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.Warning("test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"WARNING","host":"test","caller":"[^"]+:[0-9]+","message":"test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Error)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.Error("test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"ERROR","host":"test","caller":"[^"]+:[0-9]+","message":"test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Critical)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.Critical("test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"CRITICAL","host":"test","caller":"[^"]+:[0-9]+","message":"test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Alert)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.Alert("test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"ALERT","host":"test","caller":"[^"]+:[0-9]+","message":"test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Emergency)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.Emergency("test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"EMERGENCY","host":"test","caller":"[^"]+:[0-9]+","message":"test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})
}

// nolint: paralleltest
func TestLogger_Fatal(t *testing.T) {
	backup := exitFn

	t.Cleanup(func() { exitFn = backup }) // nolint: paralleltest

	exitFn = func(i int) { /* noop */ } // nolint: paralleltest

	t.Run("success(Fatal)", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		l := Must(New(buf))
		s := EMERGENCY
		l.Fatal(s, "test")
		expect := regexp.MustCompile(`^{"timestamp":"[0-9T:\.\+\-Z]+","severity":"` + l.uppercase(s) + `","caller":"[^"]+:[0-9]+","message":"test"}` + defaultLineSeparator)
		actual := buf.String()
		FailIfNotRegexpMatchString(t, expect, actual)
	})
}

func TestLogger_Panic(t *testing.T) {
	t.Parallel()

	t.Run("success(Panic)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf))
		const expectMessage = "test"
		defer func() {
			if err := recover(); err != nil {
				actualMessage, ok := err.(string)
				if !ok {
					FailIfNotEqual(t, "string", reflect.TypeOf(err))
					return // nolint: nlreturn
				}
				FailIfNotEqual(t, expectMessage, actualMessage)
			}
		}()

		s := EMERGENCY
		l.Panic(s, expectMessage)
		expect := regexp.MustCompile(`^{"timestamp":"[0-9T:\.\+\-Z]+","severity":` + l.uppercase(s) + `,"caller":"[^"]+:[0-9]+","message":` + expectMessage + `}` + defaultLineSeparator)
		actual := buf.String()
		FailIfNotRegexpMatchString(t, expect, actual)
	})
}
