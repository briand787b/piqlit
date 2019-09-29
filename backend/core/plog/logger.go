package plog

// Logger x
type Logger interface {
	Infow(msg string, keysAndValues ...interface{})
}
