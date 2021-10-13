// nolint: testpackage
package rec

import (
	"bytes"
	"reflect"
	"regexp"
	"testing"
)

// nolint: paralleltest
func TestF(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	replacer := Must(New(buf))

	backup := defaultLogger

	t.Cleanup(func() { defaultLogger = backup }) // nolint: paralleltest

	defaultLogger = replacer // nolint: paralleltest

	t.Run("success(Print)", func(t *testing.T) {
		t.Cleanup(buf.Reset)
		F().Printf(DEFAULT, "%s format", "test")
		expect := regexp.MustCompile(`^{"timestamp":"[0-9T:\.\+\-Z]+","severity":"DEFAULT","caller":"[^"]+:[0-9]+","message":"test format"}` + defaultLineSeparator)
		actual := buf.String()
		FailIfNotRegexpMatchString(t, expect, actual)
	})
}

func Test_formatLogger_Print(t *testing.T) {
	t.Parallel()

	t.Run("success(Print)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.F().Printf(DEFAULT, "printf: %s", "test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"DEFAULT","host":"test","caller":"[^"]+:[0-9]+","message":"printf: test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Default)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.F().Defaultf("printf: %s", "test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"DEFAULT","host":"test","caller":"[^"]+:[0-9]+","message":"printf: test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Debug)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.F().Debugf("printf: %s", "test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"DEBUG","host":"test","caller":"[^"]+:[0-9]+","message":"printf: test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Info)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.F().Infof("printf: %s", "test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"INFO","host":"test","caller":"[^"]+:[0-9]+","message":"printf: test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Notice)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.F().Noticef("printf: %s", "test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"NOTICE","host":"test","caller":"[^"]+:[0-9]+","message":"printf: test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Warning)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.F().Warningf("printf: %s", "test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"WARNING","host":"test","caller":"[^"]+:[0-9]+","message":"printf: test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Error)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.F().Errorf("printf: %s", "test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"ERROR","host":"test","caller":"[^"]+:[0-9]+","message":"printf: test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Critical)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.F().Criticalf("printf: %s", "test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"CRITICAL","host":"test","caller":"[^"]+:[0-9]+","message":"printf: test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Alert)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.F().Alertf("printf: %s", "test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"ALERT","host":"test","caller":"[^"]+:[0-9]+","message":"printf: test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(Emergency)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		l.F().Emergencyf("printf: %s", "test")
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"EMERGENCY","host":"test","caller":"[^"]+:[0-9]+","message":"printf: test"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
	})
}

// nolint: paralleltest
func Test_formatLogger_Fatal(t *testing.T) {
	backup := exitFn

	t.Cleanup(func() { exitFn = backup }) // nolint: paralleltest

	exitFn = func(i int) { /* noop */ } // nolint: paralleltest

	t.Run("success(Fatal)", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		s := EMERGENCY
		l.F().Fatalf(s, "printf: %s", "test")
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"` + l.uppercase(s) + `","host":"test","caller":"[^"]+:[0-9]+","message":"printf: test"}` + defaultLineSeparator)
		actual := buf.String()
		FailIfNotRegexpMatchString(t, expect, actual)
	})
}

func Test_formatLogger_Panic(t *testing.T) {
	t.Parallel()

	t.Run("success(Panic)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf))

		const expectMessage = "printf: test"
		defer func() {
			if err := recover(); err != nil {
				actualMessage, ok := err.(string)
				if !ok {
					FailIfNotEqual(t, "string", reflect.TypeOf(err))
					return
				}
				FailIfNotEqual(t, expectMessage, actualMessage)
			}
		}()

		s := EMERGENCY
		l.F().Panicf(s, "printf: %s", "test")
		expect := regexp.MustCompile(`^{"timestamp":"[0-9T:\.\+\-Z]+","severity":` + l.uppercase(s) + `,"caller":"[^"]+:[0-9]+","message":` + expectMessage + `}` + defaultLineSeparator)
		actual := buf.String()
		FailIfNotRegexpMatchString(t, expect, actual)
	})
}
