package rec

import (
	"bytes"
	"log"
	"os"
	"time"
)

type stdLogger struct {
	severity Severity
	l        *Logger
}

func (s *stdLogger) Write(b []byte) (int, error) {
	s.l.write(time.Now(), s.severity, string(bytes.TrimSpace(b)))

	return len(b), nil
}

// ReplaceStdLogger replaces the logger in Go standard log package with `*rec_Logger`
// and returns a function for rollback logger.
func ReplaceStdLogger(l *Logger, severity Severity) (rollback func()) {
	const defaultCallerSkip = 2

	backupFlags := log.Flags()
	backupPrefix := log.Prefix()

	log.SetFlags(0)
	log.SetPrefix("")
	log.SetOutput(&stdLogger{severity, l.AddCallerSkip(defaultCallerSkip)})

	return func() {
		log.SetFlags(backupFlags)
		log.SetPrefix(backupPrefix)
		log.SetOutput(os.Stderr)
	}
}
