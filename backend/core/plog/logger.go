package plog

// LogWriter writes to the concrete type that
type LogWriter interface {
	Println(v ...interface{})
}

// Logger controls the format of logs written to the underylying LogWriter
type Logger struct {
	L LogWriter
}

// NewLogger returns a new Logger holding the provided LogWriter
func NewLogger(lw LogWriter) *Logger {
	return &Logger{L: lw}
}

// Info writes INFO-lvl logs
func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.write("[INFO] ", msg, keysAndValues)
}

// Warn writes WARN-lvl logs
func (l *Logger) Warn(msg string, keysAndValues ...interface{}) {
	l.write("[WARN] ", msg, keysAndValues)
}

func (l *Logger) write(lvl, msg string, knvs ...interface{}) {
	l.L.Println(lvl, msg, ": ", knvs)
}
