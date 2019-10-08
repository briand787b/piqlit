package plog

// LogWriter writes to the concrete type that
type LogWriter interface {
	Println(v ...interface{})
}

// Logger represents anything that can format logs correctly
type Logger interface {
	// generic logging
	Error(msg string, args ...interface{})
	Info(msg string, args ...interface{})

	// specific event logging
	Invalid(subj interface{}, reason string)
}
