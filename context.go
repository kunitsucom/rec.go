package rec

import (
	"context"
)

type contextKey int

const (
	_ contextKey = iota
	key
)

// ContextLogger returns *rec.Logger that the context has.
// If the context does not have *rec.Logger, ContextLogger returns default logger.
func ContextLogger(ctx context.Context) *Logger {
	l, ok := ctx.Value(key).(*Logger)
	if !ok || l == nil {
		defaultLogger.AddCallerSkip(1).Error("rec: ContextLogger returns defaultLogger because ctx does not contain *rec.Logger")
		return defaultLogger
	}

	return l
}

// ContextWithLogger returns context.Context that has *rec.Logger.
func ContextWithLogger(parent context.Context, l *Logger) context.Context {
	return context.WithValue(parent, key, l)
}
