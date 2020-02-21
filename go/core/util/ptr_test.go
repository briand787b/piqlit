package util_test

import (
	"testing"

	"github.com/briand787b/piqlit/core/util"
)

func TestStrPtr(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name, arg string
	}{
		{"non-empty string works", "a"},
		{"empty string works", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if ret := *util.StrPtr(tt.arg); ret != tt.arg {
				t.Fatalf("expected de-referenced value to be %s, was %s", tt.arg, ret)
			}
		})
	}
}
