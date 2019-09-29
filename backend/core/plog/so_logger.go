package plog

import (
	"fmt"
	"log"
)

// SOLogger x
type SOLogger struct{}

// Infow is a placeholder and should be removed in the future
func (sol *SOLogger) Infow(msg string, keysAndValues ...interface{}) {
	log.Printf(fmt.Sprintf("%s: %v\n", msg, keysAndValues))
}
