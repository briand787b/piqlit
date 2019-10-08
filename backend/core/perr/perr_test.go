package perr_test

import (
	"database/sql"
	"testing"

	"github.com/briand787b/piqlit/core/perr"

	"github.com/pkg/errors"
)

func TestErrCause(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		err         error
		wrapErrMsgs []string
		expCause    error
	}{
		{"1_layer_in_tact", sql.ErrNoRows, []string{"a"}, sql.ErrNoRows},
		{"2_layers_in_tact", sql.ErrNoRows, []string{"a", "b"}, sql.ErrNoRows},
		{"3_layers_in_tact", sql.ErrNoRows, []string{"a", "b", "c"}, sql.ErrNoRows},
		{"empty_val_in_tact", perr.NewErrInvalid(""), []string{"a"}, perr.ErrInvalid},
		{"full_val_in_tact", perr.NewErrInvalid("z"), []string{"a"}, perr.ErrInvalid},
		{"2_layer_full_val_in_tact", perr.NewErrInvalid("z"), []string{"a", "b"}, perr.ErrInvalid},
		{"empty_not_found_in_tact", perr.NewErrNotFound(nil), []string{"a"}, perr.ErrNotFound},
		{"full_not_found_in_tact", perr.NewErrNotFound(sql.ErrNoRows), []string{"a"}, perr.ErrNotFound},
		{"2_layer_not_found_in_tact", perr.NewErrNotFound(sql.ErrNoRows), []string{"a", "b"}, perr.ErrNotFound},
		{"empty_internal_in_tact", perr.NewErrInternal(nil), []string{"a"}, perr.ErrInternal},
		{"full_internal_in_tact", perr.NewErrInternal(sql.ErrNoRows), []string{"a"}, perr.ErrInternal},
		{"2_layer_internal_in_tact", perr.NewErrInternal(sql.ErrNoRows), []string{"a", "b"}, perr.ErrInternal},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var err error
			for _, msg := range tt.wrapErrMsgs {
				err = errors.Wrap(tt.err, msg)
			}

			if retErr := errors.Cause(err); tt.expCause != retErr {
				t.Fatalf("expected error to be %v, was %v", tt.expCause, retErr)
			}
		})
	}
}

func TestGetExternalMgs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		es     []error
		expMsg string
	}{
		{"nil_err_returns_empty_msg", nil, ""},
		{"1_err_returns_its_msg", []error{errors.New("a")}, "a"},
		{"2_errs_returns_1st_msg", []error{errors.New("a"), errors.New("b")}, "a"},
		{"3_errs_returns_1st_msg", []error{errors.New("a"), errors.New("b"), errors.New("c")}, "a"},
		{"base_val_returns_cause_msg", []error{perr.ErrInvalid}, "request invalid"},
		{"val_returns_internal+provided_msg", []error{perr.NewErrInvalid("a")}, "request invalid: a"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var err error
			if tt.es != nil || len(tt.es) > 0 {
				err = tt.es[0]
				for i := 1; i < len(tt.es); i++ {
					err = errors.Wrap(err, tt.es[i].Error())
				}
			}

			if retMsg := perr.GetExternalMsg(err); tt.expMsg != retMsg {
				t.Fatalf("expected msg to be %s, was %s", tt.expMsg, retMsg)
			}
		})
	}
}

func TestGetInternallMgs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		es     []error
		expMsg string
	}{
		{"nil_err_returns_empty_msg", nil, ""},
		{"1_err_returns_its_msg", []error{errors.New("a")}, "a"},
		{"2_errs_returns_1st_msg", []error{errors.New("a"), errors.New("b")}, "b: a"},
		{"3_errs_returns_1st_msg", []error{errors.New("a"), errors.New("b"), errors.New("c")}, "c: b: a"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var err error
			if tt.es != nil || len(tt.es) > 0 {
				err = tt.es[0]
				for i := 1; i < len(tt.es); i++ {
					err = errors.Wrap(err, tt.es[i].Error())
				}
			}

			if retMsg := perr.GetInternalMsg(err); tt.expMsg != retMsg {
				t.Fatalf("expected msg to be %s, was %s", tt.expMsg, retMsg)
			}
		})
	}
}
