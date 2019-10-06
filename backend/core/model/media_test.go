package model_test

import (
	"testing"

	"github.com/briand787b/piqlit/core/model"
	_ "github.com/briand787b/piqlit/core/test"
)

func TestMediaValidate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		m      model.Media
		expErr error
	}{
		{},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// tt.m.Validate()
		})
	}
}
