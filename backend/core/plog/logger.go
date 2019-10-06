package plog

import (
	"encoding/json"
	"fmt"
)

// LogWriter writes to the concrete type that
type LogWriter interface {
	Printf(v ...interface{})
}

// Logger controls the format of logs written to the underylying LogWriter
type Logger struct {
	l LogWriter
}

// NewLogger returns a new Logger holding the provided LogWriter
func NewLogger(lw LogWriter) *Logger {
	return &Logger{l: lw}
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.write("[ERROR]", msg, args...)
}

// // Info writes INFO-lvl logs
// func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
// 	l.write("[INFO] ", msg, keysAndValues)
// }

func (l *Logger) Invalid(subj interface{}, reason string) {
	l.write("[INVALID]", fmt.Sprintf("%T failed validation", subj),
		[]string{"reason", reason},
	)
}

// // Warn writes WARN-lvl logs
// func (l *Logger) Warn(msg string, keysAndValues ...interface{}) {
// 	l.write("[WARN] ", msg, keysAndValues)
// }

func (l *Logger) write(lvl, msg string, kvs ...interface{}) {
	if kvs != nil && len(kvs)%2 != 0 {
		l.Error("uneven number of args provided to Logger",
			"number_of_args_provided", len(kvs),
		)
		l.l.Printf("%s %s: %v", lvl, msg, map[string]interface{}{
			"Logger_Args": kvs,
			"State":       "invalid",
		})
		return
	}

	m := make(map[string]interface{})
	for i := 0; i < len(kvs); i += 2 {
		key, ok := (kvs[i]).(string)
		if !ok {
			l.Error("Logger key not assertable to string", "key", kvs[i])
			key = fmt.Sprintf("%v", kvs[i])
		}

		m[key] = kvs[i+1]
	}

	bs, err := json.Marshal(m)
	if err != nil {
		l.Error("could not marshal key values to JSON", "kvs", kvs)
	}

	l.l.Printf("%s %s: %v", lvl, msg, string(bs))
}
