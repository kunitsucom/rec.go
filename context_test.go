package rec_test

import (
	"context"
	"io"
	"testing"

	"github.com/rec-logger/rec.go"
)

func TestContextGet(t *testing.T) {
	t.Parallel()

	testLogger := rec.L().RenewWriter(io.Discard)
	tests := []struct {
		name   string
		ctx    context.Context
		expect *rec.Logger
	}{
		{"success()", rec.ContextWithLogger(context.Background(), testLogger), testLogger},
		{"success(defaultLogger)", context.Background(), rec.L()},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := rec.ContextLogger(tt.ctx)
			rec.FailIfNotEqual(t, tt.expect, actual)
		})
	}
}
