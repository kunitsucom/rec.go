package rec

import (
	"log"
	"os"
)

const callerSkipForStdLogger = 2

func NewStdLogger(l *Logger, severity Severity) *log.Logger {
	return log.New(Must(l.Renew(WithDefaultSeverity(severity))).AddCallerSkip(callerSkipForStdLogger), "", 0)
}

// ReplaceStdLogger replaces the logger in Go standard log package with `*rec_Logger`
// and returns a function for rollback logger.
func ReplaceStdLogger(l *Logger, severity Severity) (rollback func()) {
	backupFlags := log.Flags()
	backupPrefix := log.Prefix()

	log.SetFlags(0)
	log.SetPrefix("")
	log.SetOutput(Must(l.Renew(WithDefaultSeverity(severity))).AddCallerSkip(callerSkipForStdLogger))

	return func() {
		log.SetFlags(backupFlags)
		log.SetPrefix(backupPrefix)
		log.SetOutput(os.Stderr)
	}
}
