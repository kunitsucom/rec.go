// nolint: testpackage
package rec

import (
	"bytes"
	"reflect"
	"regexp"
	"testing"
)

// nolint: paralleltest
func TestE(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	replacer := Must(New(buf))

	backup := defaultLogger

	t.Cleanup(func() { defaultLogger = backup }) // nolint: paralleltest

	defaultLogger = replacer // nolint: paralleltest

	t.Run("success(Print)", func(t *testing.T) {
		t.Cleanup(buf.Reset)
		E().Print(DEFAULT, errForTest)
		expect := regexp.MustCompile(`^{"timestamp":"[0-9T:\.\+\-Z]+","severity":"DEFAULT","caller":"[^"]+:[0-9]+","message":"test error","error":"test error","errorStacktrace":"test error"}` + defaultLineSeparator)
		actual := buf.String()
		FailIfNotRegexpMatchString(t, expect, actual)
	})
}

func Test_errorLogger_Print(t *testing.T) {
	t.Parallel()

	t.Run("success(PrintErr)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Print(DEFAULT, testFormatErrorValue, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"DEFAULT","host":"test","caller":"[^"]+:[0-9]+","message":"EOF","error":"EOF","errorStacktrace":"-Inf","test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, testFormatErrorValue, err)
	})

	t.Run("success(PrintErr,nil)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Print(DEFAULT, nil, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"DEFAULT","host":"test","caller":"[^"]+:[0-9]+","message":"<nil>","error":null,"errorStacktrace":null,"test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, nil, err)
	})

	t.Run("success(DefaultErr)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Default(testFormatErrorValue, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"DEFAULT","host":"test","caller":"[^"]+:[0-9]+","message":"EOF","error":"EOF","errorStacktrace":"-Inf","test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, testFormatErrorValue, err)
	})

	t.Run("success(DefaultErr,nil)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Default(nil, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"DEFAULT","host":"test","caller":"[^"]+:[0-9]+","message":"<nil>","error":null,"errorStacktrace":null,"test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, nil, err)
	})

	t.Run("success(DebugErr)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Debug(testFormatErrorValue, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"DEBUG","host":"test","caller":"[^"]+:[0-9]+","message":"EOF","error":"EOF","errorStacktrace":"-Inf","test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, testFormatErrorValue, err)
	})

	t.Run("success(DebugErr,nil)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Debug(nil, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"DEBUG","host":"test","caller":"[^"]+:[0-9]+","message":"<nil>","error":null,"errorStacktrace":null,"test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, nil, err)
	})

	t.Run("success(InfoErr)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Info(testFormatErrorValue, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"INFO","host":"test","caller":"[^"]+:[0-9]+","message":"EOF","error":"EOF","errorStacktrace":"-Inf","test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, testFormatErrorValue, err)
	})

	t.Run("success(InfoErr,nil)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Info(nil, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"INFO","host":"test","caller":"[^"]+:[0-9]+","message":"<nil>","error":null,"errorStacktrace":null,"test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, nil, err)
	})

	t.Run("success(NoticeErr)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Notice(testFormatErrorValue, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"NOTICE","host":"test","caller":"[^"]+:[0-9]+","message":"EOF","error":"EOF","errorStacktrace":"-Inf","test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, testFormatErrorValue, err)
	})

	t.Run("success(NoticeErr,nil)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Notice(nil, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"NOTICE","host":"test","caller":"[^"]+:[0-9]+","message":"<nil>","error":null,"errorStacktrace":null,"test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, nil, err)
	})

	t.Run("success(WarningErr)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Warning(testFormatErrorValue, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"WARNING","host":"test","caller":"[^"]+:[0-9]+","message":"EOF","error":"EOF","errorStacktrace":"-Inf","test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, testFormatErrorValue, err)
	})

	t.Run("success(WarningErr,nil)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Warning(nil, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"WARNING","host":"test","caller":"[^"]+:[0-9]+","message":"<nil>","error":null,"errorStacktrace":null,"test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, nil, err)
	})

	t.Run("success(ErrorErr)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Error(testFormatErrorValue, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"ERROR","host":"test","caller":"[^"]+:[0-9]+","message":"EOF","error":"EOF","errorStacktrace":"-Inf","test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, testFormatErrorValue, err)
	})

	t.Run("success(ErrorErr,nil)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Error(nil, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"ERROR","host":"test","caller":"[^"]+:[0-9]+","message":"<nil>","error":null,"errorStacktrace":null,"test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, nil, err)
	})

	t.Run("success(CriticalErr)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Critical(testFormatErrorValue, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"CRITICAL","host":"test","caller":"[^"]+:[0-9]+","message":"EOF","error":"EOF","errorStacktrace":"-Inf","test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, testFormatErrorValue, err)
	})

	t.Run("success(CriticalErr,nil)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Critical(nil, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"CRITICAL","host":"test","caller":"[^"]+:[0-9]+","message":"<nil>","error":null,"errorStacktrace":null,"test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, nil, err)
	})

	t.Run("success(AlertErr)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Alert(testFormatErrorValue, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"ALERT","host":"test","caller":"[^"]+:[0-9]+","message":"EOF","error":"EOF","errorStacktrace":"-Inf","test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, testFormatErrorValue, err)
	})

	t.Run("success(AlertErr,nil)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Alert(nil, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"ALERT","host":"test","caller":"[^"]+:[0-9]+","message":"<nil>","error":null,"errorStacktrace":null,"test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, nil, err)
	})

	t.Run("success(EmergencyErr)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Emergency(testFormatErrorValue, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"EMERGENCY","host":"test","caller":"[^"]+:[0-9]+","message":"EOF","error":"EOF","errorStacktrace":"-Inf","test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, testFormatErrorValue, err)
	})

	t.Run("success(EmergencyErr,nil)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf, WithUseHostnameField(true), WithHostnameFieldKey("host"), WithHostnameFieldValue("test")))
		t.Cleanup(buf.Reset)
		err := l.E().Emergency(nil, String("test", "field")).Err()
		actual := buf.String()
		expect := regexp.MustCompile(`^{"timestamp":"[0-9]+\-[0-9]+\-[0-9]+T[0-9]+:[0-9]+:[0-9]+[\.0-9]*[0-9Z:\+\-]+","severity":"EMERGENCY","host":"test","caller":"[^"]+:[0-9]+","message":"<nil>","error":null,"errorStacktrace":null,"test":"field"}` + defaultLineSeparator)
		FailIfNotRegexpMatchString(t, expect, actual)
		FailIfNotErrorIs(t, nil, err)
	})
}

// nolint: paralleltest
func Test_errorLogger_Fatal(t *testing.T) {
	backup := exitFn

	t.Cleanup(func() { exitFn = backup }) // nolint: paralleltest

	exitFn = func(i int) { /* noop */ } // nolint: paralleltest

	t.Run("success(FatalErr)", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		l := Must(New(buf))
		t.Cleanup(buf.Reset)
		s := EMERGENCY
		l.E().Fatal(s, testFormatErrorValue, String("test", "field"))
		expect := regexp.MustCompile(`^{"timestamp":"[0-9T:\.\+\-Z]+","severity":"` + l.uppercase(s) + `","caller":"[^"]+:[0-9]+","message":"EOF","error":"EOF","errorStacktrace":"-Inf","test":"field"}` + defaultLineSeparator)
		actual := buf.String()
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(FatalErr,nil)", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		l := Must(New(buf))
		t.Cleanup(buf.Reset)
		s := EMERGENCY
		l.E().Fatal(s, nil, String("test", "field"))
		expect := regexp.MustCompile(`^{"timestamp":"[0-9T:\.\+\-Z]+","severity":"` + l.uppercase(s) + `","caller":"[^"]+:[0-9]+","message":"<nil>","error":null,"errorStacktrace":null,"test":"field"}` + defaultLineSeparator)
		actual := buf.String()
		FailIfNotRegexpMatchString(t, expect, actual)
	})
}

func Test_errorLogger_Panic(t *testing.T) {
	t.Parallel()

	t.Run("success(PanicErr)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf))
		t.Cleanup(buf.Reset)
		expectError := testFormatErrorValue
		defer func() {
			if err := recover(); err != nil {
				actualError, ok := err.(error)
				if !ok {
					FailIfNotEqual(t, "error", reflect.TypeOf(err))
					return
				}
				FailIfNotErrorIs(t, expectError, actualError)
			}
		}()

		s := EMERGENCY
		l.E().Panic(s, testFormatErrorValue, String("test", "field"))
		expect := regexp.MustCompile(`^{"timestamp":"[0-9T:\.\+\-Z]+","severity":` + l.uppercase(s) + `,"caller":"[^"]+:[0-9]+","message":"-Inf"}` + defaultLineSeparator)
		actual := buf.String()
		FailIfNotRegexpMatchString(t, expect, actual)
	})

	t.Run("success(PanicErr,nil)", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(nil)
		l := Must(New(buf))
		t.Cleanup(buf.Reset)
		defer func() {
			if err := recover(); err != nil {
				actual, ok := err.(error)
				if !ok {
					FailIfNotEqual(t, "error", reflect.TypeOf(err))
					return
				}
				FailIfNotEqual(t, "<nil>", actual.Error())
			}
		}()

		s := EMERGENCY
		l.E().Panic(s, nil, String("test", "field"))
		expect := regexp.MustCompile(`^{"timestamp":"[0-9T:\.\+\-Z]+","severity":` + l.uppercase(s) + `,"caller":"[^"]+:[0-9]+","message":"<nil>"}` + defaultLineSeparator)
		actual := buf.String()
		FailIfNotRegexpMatchString(t, expect, actual)
	})
}
