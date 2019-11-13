package perr_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/briand787b/piqlit/core/perr"
	"github.com/briand787b/piqlit/core/plog/plogtest"

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
		err    error
		expMsg string
	}{
		{"nil_err_returns_empty_msg", nil, ""},
		{"unknown_err_returns_internal_server_err", errors.New("a"), "Internal Server Error"},
		{"auth_err_returns_unauth_msg", perr.ErrUnauthorized, "Request Not Authorized to Perform Action"},
		{"invalid_err_returns_internal_server_err", perr.NewErrInvalid("b"), "request invalid: b"},
		{"uninit_invalid_err_returns_internal_server_err", perr.ErrInvalid, "Internal Server Error"},
		{"not_found_err_returns_invalid_msg", perr.ErrNotFound, "Resource Not Found"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			if retMsg := perr.GetExternalMsg(ctx, plogtest.NewMockLogger(), tt.err); tt.expMsg != retMsg {
				t.Fatalf("expected msg to be %s, was %s", tt.expMsg, retMsg)
			}
		})
	}
}
