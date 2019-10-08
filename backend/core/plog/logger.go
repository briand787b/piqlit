package plog

import "context"

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
	Invalid(ctx context.Context, subj interface{}, reason string)
	Query(ctx context.Context, qry string, args ...interface{})
}
