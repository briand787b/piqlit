package plogtest

// SpyLogWriter is a spying implementation of plog.Logger
type SpyLogWriter struct {
	PrintlnCallCount int
	PrintlnArgs      []interface{}
}

// Println is the spied implementation of plog.LogWriter.Println
func (s *SpyLogWriter) Println(v ...interface{}) {
	defer func() { s.PrintlnCallCount++ }()
	s.PrintlnArgs = append(s.PrintlnArgs, v)
}
