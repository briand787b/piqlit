package plogtest

import (
	"context"

	"github.com/briand787b/piqlit/core/plog"
)

const (
	// SpanID is the dummy SpanID - only use for tests
	SpanID = "SpanID"

	// TraceID is the dummy TraceID - only use for tests
	TraceID = "TraceID"
)

// SpannedTracedCtx is a helper function for creating a context that doesn't
// log `SpanID not found` or `TraceID not found` errors when running tests
func SpannedTracedCtx() context.Context {
	return plog.StoreSpanIDTraceID(context.Background(), SpanID, TraceID)
}
