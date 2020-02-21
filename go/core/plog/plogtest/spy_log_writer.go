package plogtest

import "github.com/briand787b/piqlit/core/plog"

var _ plog.LogWriter = &SpyLogWriter{}

// SpyLogWriter is a spying implementation of plog.LogWriter
type SpyLogWriter struct {
	PrintlnCallCount int
	PrintlnArgs      [][]interface{}
}

// Println is the spied implementation of plog.LogWriter.Println
func (s *SpyLogWriter) Println(v ...interface{}) {
	defer func() { s.PrintlnCallCount++ }()
	s.PrintlnArgs = append(s.PrintlnArgs, v)
}
