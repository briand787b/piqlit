package plog

import "context"

// spanIDCtxKeyType is the type that keys the SpanID in the context
type spanIDCtxKeyType int

// traceIDCtxKeyType is the type that keys the TraceID in the context
type traceIDCtxKeyType int

const (
	// spanIDCtxKey is the only valid value for SpanIDCtxKeyType
	spanIDCtxKey spanIDCtxKeyType = 0

	// traceIDCtxKey is the only valid value for TraceIDCtxKeyType
	traceIDCtxKey traceIDCtxKeyType = 0
)

// StoreSpanID returns a context with a stored spanID
func StoreSpanID(ctx context.Context, spanID string) context.Context {
	return context.WithValue(ctx, spanIDCtxKey, spanID)
}

// StoreTraceID returns a context with a stored traceID
func StoreTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDCtxKey, traceID)
}

// StoreSpanIDTraceID store both SpanID and TraceID
func StoreSpanIDTraceID(ctx context.Context, spanID, traceID string) context.Context {
	return StoreTraceID(StoreSpanID(ctx, spanID), traceID)
}

func getSpanID(ctx context.Context) string {
	sID, _ := ctx.Value(spanIDCtxKey).(string)
	return sID
}

func getTraceID(ctx context.Context) string {
	tID, _ := ctx.Value(traceIDCtxKey).(string)
	return tID
}
