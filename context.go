package rec

import (
	"context"
)

type contextKeyType struct{}

var contextKey *contextKeyType // nolint: gochecknoglobals

// ContextLogger returns *rec.Logger that the context has.
// If the context does not have *rec.Logger, ContextLogger returns default logger.
func ContextLogger(ctx context.Context) *Logger {
	l, ok := ctx.Value(contextKey).(*Logger)
	if !ok || l == nil {
		return defaultLogger
	}

	return l
}

// ContextWithLogger returns context.Context that has *rec.Logger.
func ContextWithLogger(parent context.Context, l *Logger) context.Context {
	return context.WithValue(parent, contextKey, l)
}
