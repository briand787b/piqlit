package test

import (
	"log"
	"testing"
	"time"

	"github.com/briand787b/piqlit/core/plog"
)

// Helper is a testing utility that facilitates rolling back
// whatever changes were made to persistent storage during
// the course of the test.  It is meant to be embedded by
// implementation specific adaptors
type Helper struct {
	T  *testing.T
	L  plog.Logger
	Tm time.Time
	CF *CleanFunc
}

// Clean cleans up after the test.  This function is
// typically deferred after all test setup operations
// are complete
func (h *Helper) Clean() {
	(*h.CF)()
}

// CleanFunc is a function that takes no arguments and
// returns no values.  It is intended to represent
// any function that serves to clean up after
// running some test(s)
//
// The most valuable feature of the CleanFunc is that
// additional operations can be added to it through
// a pointer method receiver.  This means that only
// one call to defer is required to defer every
// operation that needs to be cleaned after its
// initial calling
type CleanFunc func()

// NewCleaner returns a pointer to a cleaner,
// taking in a bare function
func NewCleaner(f func()) *CleanFunc {
	c := CleanFunc(f)
	return &c
}

// Add adds a new operation to the internals of
// the cleaning function method receiver
//
// The provided function is executed BEFORE
// the existing operations.  The order of
// execution from the Cleaners perspective
// can be thought of as a stack
func (c *CleanFunc) Add(f func()) {
	oldFn := *c
	newFn := func() {
		f()
		oldFn()
	}

	*c = CleanFunc(newFn)
}

// SetTimeout sets a timeout for the test, which is avoided if
// a value is sent on the returned channel
func SetTimeout(t *testing.T, to time.Duration) chan<- struct{} {
	ch := make(chan struct{})
	go func(t *testing.T, to time.Duration, c <-chan struct{}) {
		tm := time.NewTimer(to)
		select {
		case <-c:
			if !tm.Stop() {
				<-tm.C
			}
		case <-tm.C:
			log.Fatalf("Helper init exceeded timeout of %v\n", to)
		}
	}(t, to, ch)

	return ch
}
