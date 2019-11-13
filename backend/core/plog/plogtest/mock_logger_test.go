package plogtest_test

import (
	"context"
	"strconv"
	"sync"
	"testing"

	"github.com/briand787b/piqlit/core/plog/plogtest"
)

func TestConcurrentLoggerWrites(t *testing.T) {
	l := plogtest.MockLogger{}
	ctx := context.Background()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			is := strconv.Itoa(i)
			wg.Add(1)
			go func() { l.Close(ctx, nil); wg.Done() }()

			wg.Add(1)
			go func() { l.Error(ctx, is); wg.Done() }()

			wg.Add(1)
			go func() { l.Info(ctx, is); wg.Done() }()

			wg.Add(1)
			go func() { l.Invalid(ctx, is, is); wg.Done() }()

			wg.Add(1)
			go func() { l.Query(ctx, is); wg.Done() }()

			wg.Done()
		}(i)
	}
	wg.Wait()
}
