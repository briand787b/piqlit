package plog

import (
	"context"
	"io"
)

// LogWriter writes to the concrete type that
type LogWriter interface {
	Println(v ...interface{})
}

// Logger represents anything that can format logs correctly
type Logger interface {
	// generic logging
	Error(ctx context.Context, msg string, args ...interface{})
	Info(ctx context.Context, msg string, args ...interface{})

	// specific event logging
	Close(ctx context.Context, c io.Closer)
	Invalid(ctx context.Context, subj interface{}, reason string)
	Query(ctx context.Context, qry string, args ...interface{})
}
