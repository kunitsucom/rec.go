// nolint: testpackage
package rec

import (
	"bytes"
	"log"
	"regexp"
	"runtime"
	"strconv"
	"testing"
)

func TestReplaceStdLogger(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	l := Must(New(buf))

	tests := []struct {
		name     string
		l        *Logger
		severity Severity
	}{
		{"test", l, ERROR},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rollback := ReplaceStdLogger(tt.l, tt.severity)
			defer rollback()

			log.Printf("test")                    // <-
			_, _, linenum, _ := runtime.Caller(0) // <-
			linenum--                             // <- Get the number of lines executed by `log.Printf()`.

			expect := regexp.MustCompile(`^{"timestamp":"[0-9T:\.\+\-Z]+","severity":"` + l.uppercase(tt.severity) + `","caller":"[^"]+:` + strconv.Itoa(linenum) + `","message":"test"}` + defaultLineSeparator)
			actual := buf.String()
			FailIfNotRegexpMatchString(t, expect, actual)
		})
	}
}
