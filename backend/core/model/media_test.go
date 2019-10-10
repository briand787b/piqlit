package model_test

import (
	"testing"

	"github.com/briand787b/piqlit/core/model"
	"github.com/briand787b/piqlit/core/model/modeltest"
	"github.com/briand787b/piqlit/core/obj"
	"github.com/briand787b/piqlit/core/perr"
	"github.com/briand787b/piqlit/core/plog/plogtest"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

func TestMediaValidate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		m          model.Media
		expBaseErr error
		expReason  string
	}{
		{
			"media_with_full_fields_passes",
			model.Media{
				Name:         "full_name",
				Length:       1,
				Encoding:     obj.GIF,
				UploadStatus: obj.UploadDone,
			},
			nil,
			"",
		},
		{
			"media_with_empty_name_fails",
			model.Media{
				Length:       1,
				Encoding:     obj.GIF,
				UploadStatus: obj.UploadDone,
			},
			perr.ErrInvalid,
			"empty field: name",
		},
		{
			"media_with_zero_length_fails",
			model.Media{
				Name:         "full_name",
				Length:       0,
				Encoding:     obj.GIF,
				UploadStatus: obj.UploadDone},
			perr.ErrInvalid,
			"non-positive length",
		},
		{
			"media_with_negative_length_fails",
			model.Media{
				Name:         "full_name",
				Length:       -1,
				Encoding:     obj.GIF,
				UploadStatus: obj.UploadDone},
			perr.ErrInvalid,
			"non-positive length",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ml := plogtest.MockLogger{}
			err := tt.m.Validate(plogtest.SpannedTracedCtx(), &ml)

			t.Logf("returned error: %s", err)
			retErrCause := errors.Cause(err)
			t.Logf("returned base error: %s", retErrCause)

			t.Logf("expected error: %s", tt.expBaseErr)

			if !cmp.Equal(tt.expBaseErr, retErrCause) {
				t.Fatal("expected error not equal to returned error: ",
					cmp.Diff(tt.expBaseErr, retErrCause),
				)
			}

			if tt.expBaseErr == nil {
				return
			}

			if _, ok := ml.InvalidArgSub[0].(model.Media); !ok {
				t.Fatalf("expected log subj to be model.Media, was of type %T",
					ml.InvalidArgSub[0],
				)
			}

			if tt.expReason != ml.InvalidArgReason[0] {
				t.Fatalf("expected reason to be '%s', was '%s'",
					tt.expReason, ml.InvalidArgReason[0],
				)
			}
		})
	}
}

func TestMediaPersist(t *testing.T) {
	tests := []struct {
		name              string
		m                 model.Media
		msUpdateErr       []error
		msInsertErr       []error
		msAssociateErr    []error
		msUpdateExpMedia  []*model.Media
		msInsertExpMedia  []*model.Media
		msAssociateExpPID []int
		msAssociateExpCID [][]int
		expErrToBeNil     bool
	}{
		{},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := plogtest.SpannedTracedCtx()
			ml := plogtest.MockLogger{}

			ms := modeltest.MediaMockStore{}
		})
	}
}
