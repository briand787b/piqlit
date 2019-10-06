package plog

// Logger represents anything that can format logs correctly
type Logger interface {
	Error(err error, args ...interface{})
	ErrorStr(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Invalid(subj interface{}, reason string)
}
