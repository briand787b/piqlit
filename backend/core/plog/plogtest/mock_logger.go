package plogtest

import (
	"context"

	"github.com/briand787b/piqlit/core/plog"
)

var _ plog.Logger = &MockLogger{}

// MockLogger is a mocked implementation of plog.Logger
type MockLogger struct {
	ErrorCallCount int
	ErrorArgMsg    []string
	ErrorArgArgs   [][]interface{}

	InfoCallCount int
	InfoArgMsg    []string
	InfoArgArgs   [][]interface{}

	InvalidCallCount int
	InvalidArgSub    []interface{}
	InvalidArgReason []string

	QueryCallCount int
	QueryArgQry    []string
	QueryArgArgs   [][]interface{}
}

// Error x
func (m *MockLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	defer func() { m.ErrorCallCount++ }()
	m.ErrorArgMsg = append(m.ErrorArgMsg, msg)
	m.ErrorArgArgs = append(m.ErrorArgArgs, args)
}

// Info x
func (m *MockLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	defer func() { m.InfoCallCount++ }()
	m.InfoArgMsg = append(m.InfoArgMsg, msg)
	m.InfoArgArgs = append(m.InfoArgArgs, args)
}

// Invalid x
func (m *MockLogger) Invalid(ctx context.Context, subj interface{}, reason string) {
	defer func() { m.InvalidCallCount++ }()
	m.InvalidArgSub = append(m.InvalidArgSub, subj)
	m.InvalidArgReason = append(m.InvalidArgReason, reason)
}

// Query x
func (m *MockLogger) Query(ctx context.Context, qry string, args ...interface{}) {
	defer func() { m.QueryCallCount++ }()
	m.QueryArgQry = append(m.QueryArgQry, qry)
	m.QueryArgArgs = append(m.QueryArgArgs, args)
}
