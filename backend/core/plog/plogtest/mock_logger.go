package plogtest

import "github.com/briand787b/piqlit/core/plog"

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
}

// Error x
func (m *MockLogger) Error(msg string, args ...interface{}) {
	defer func() { m.ErrorCallCount++ }()
	m.ErrorArgMsg = append(m.ErrorArgMsg, msg)
	m.ErrorArgArgs = append(m.ErrorArgArgs, args)
}

// Info x
func (m *MockLogger) Info(msg string, args ...interface{}) {
	defer func() { m.InfoCallCount++ }()
	m.InfoArgMsg = append(m.InfoArgMsg, msg)
	m.InfoArgArgs = append(m.InfoArgArgs, args)
}

// Invalid x
func (m *MockLogger) Invalid(subj interface{}, reason string) {
	defer func() { m.InvalidCallCount++ }()
	m.InvalidArgSub = append(m.InvalidArgSub, subj)
	m.InvalidArgReason = append(m.InvalidArgReason, reason)
}