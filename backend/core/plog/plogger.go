package plog

import (
	"encoding/json"
	"fmt"
)

// LogWriter writes to the concrete type that
type LogWriter interface {
	Printf(v ...interface{})
}

// PLogger controls the format of logs written to the underylying LogWriter
type PLogger struct {
	l LogWriter
}

// NewPLogger returns a new PLogger holding the provided LogWriter
func NewPLogger(lw LogWriter) *PLogger {
	return &PLogger{l: lw}
}

// Error writes ERROR-lvl logs
func (l *PLogger) Error(err error, args ...interface{}) {
	l.write("[ERROR]", err.Error(), args...)
}

// ErrorStr writes ERROR-lvl logs, but takes input as a string
func (l *PLogger) ErrorStr(msg string, args ...interface{}) {
	l.write("[ERROR]", msg, args...)
}

// Info writes INFO-lvl logs
func (l *PLogger) Info(msg string, args ...interface{}) {
	l.write("[INFO] ", msg, args...)
}

// Invalid writes logs for failed validation events
func (l *PLogger) Invalid(subj interface{}, reason string) {
	l.write("[INVALID]", fmt.Sprintf("%T failed validation", subj),
		[]string{"reason", reason},
	)
}

// Warn writes WARN-lvl logs
func (l *PLogger) Warn(msg string, args ...interface{}) {
	l.write("[WARN] ", msg, args...)
}

func (l *PLogger) write(lvl, msg string, kvs ...interface{}) {
	if kvs != nil && len(kvs)%2 != 0 {
		l.ErrorStr("uneven number of args provided to PLogger",
			"number_of_args_provided", len(kvs),
		)
		l.l.Printf("%s %s: %v", lvl, msg, map[string]interface{}{
			"PLogger_Args": kvs,
			"State":        "invalid",
		})
		return
	}

	m := make(map[string]interface{})
	for i := 0; i < len(kvs); i += 2 {
		key, ok := (kvs[i]).(string)
		if !ok {
			l.ErrorStr("PLogger key not assertable to string", "key", kvs[i])
			key = fmt.Sprintf("%v", kvs[i])
		}

		m[key] = kvs[i+1]
	}

	bs, err := json.Marshal(m)
	if err != nil {
		l.ErrorStr("could not marshal key values to JSON", "kvs", kvs)
	}

	l.l.Printf("%s %s: %v", lvl, msg, string(bs))
}
