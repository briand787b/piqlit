package model_test

import (
	"testing"
	"time"

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
			"media_with_zero_length_passes",
			model.Media{
				Name:         "full_name",
				Length:       0,
				Encoding:     obj.GIF,
				UploadStatus: obj.UploadDone},
			nil,
			"",
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
	now := time.Now().UTC().Truncate(time.Second)
	tests := []struct {
		name              string
		m                 model.Media
		msGetByNameMedia  []*model.Media
		msGetByNameErr    []error
		msUpdateErr       []error
		msInsertErr       []error
		msAssociateErr    []error
		msUpdateExpMedia  []*model.Media
		msInsertExpMedia  []*model.Media
		msAssociateExpPID []int
		msAssociateExpCID [][]int
		expErrToBeNil     bool
	}{
		{
			"inserts_valid_media_with_id_0_successfully_no_error",
			model.Media{ID: 0, Name: "kiwi"},
			[]*model.Media{nil},
			[]error{errors.New("blah")},
			nil,
			[]error{nil},
			nil,
			nil,
			[]*model.Media{&model.Media{ID: 0, Name: "kiwi", CreatedAt: now, UpdatedAt: now}},
			nil,
			nil,
			true,
		},
		{
			"inserts_valid_media_with_id_0_fails_returns_error",
			model.Media{ID: 0, Name: "kiwi"},
			[]*model.Media{nil},
			[]error{errors.New("blah")},
			nil,
			[]error{errors.New("blah")},
			nil,
			nil,
			[]*model.Media{&model.Media{ID: 0, Name: "kiwi", CreatedAt: now, UpdatedAt: now}},
			nil,
			nil,
			false,
		},
		{
			"inserts_media_with_duplicate_name_returns_error",
			model.Media{ID: 0, Name: "kiwi"},
			[]*model.Media{&model.Media{}},
			[]error{nil},
			nil,
			[]error{nil},
			nil,
			nil,
			[]*model.Media{&model.Media{ID: 0, Name: "kiwi", CreatedAt: now, UpdatedAt: now}},
			nil,
			nil,
			false,
		},
		{
			"updates_valid_media_with_id_1_successfully",
			model.Media{ID: 1, Name: "peach"},
			[]*model.Media{nil},
			[]error{errors.New("blah")},
			[]error{nil},
			nil,
			nil,
			[]*model.Media{&model.Media{ID: 1, Name: "peach", UpdatedAt: now}},
			nil,
			nil,
			nil,
			true,
		},
		{
			"updates_valid_media_with_id_1_fails_returns_error",
			model.Media{ID: 1, Name: "peach"},
			[]*model.Media{nil},
			[]error{errors.New("blah")},
			[]error{errors.New("blah")},
			nil,
			nil,
			[]*model.Media{&model.Media{ID: 1, Name: "peach", UpdatedAt: now}},
			nil,
			nil,
			nil,
			false,
		},
		{
			"updates_media_with_duplicate_name_diff_id_fails",
			model.Media{ID: 1, Name: "peach"},
			[]*model.Media{&model.Media{}},
			[]error{nil},
			[]error{nil},
			nil,
			nil,
			[]*model.Media{&model.Media{ID: 1, Name: "peach", UpdatedAt: now}},
			nil,
			nil,
			nil,
			false,
		},
		{
			"updates_media_with_duplicate_name_same_id_successfully",
			model.Media{ID: 1, Name: "peach"},
			[]*model.Media{&model.Media{ID: 1}},
			[]error{nil},
			[]error{nil},
			nil,
			nil,
			[]*model.Media{&model.Media{ID: 1, Name: "peach", UpdatedAt: now}},
			nil,
			nil,
			nil,
			true,
		},
		{
			"inserts_parent_and_child_successfully",
			model.Media{ID: 0, Name: "mango", Children: []model.Media{model.Media{ID: 0, Name: "plumb"}}},
			[]*model.Media{nil, nil},
			[]error{errors.New("blah"), errors.New("blah")},
			nil,
			[]error{nil, nil},
			[]error{nil},
			nil,
			[]*model.Media{
				&model.Media{ID: 0, Name: "mango", CreatedAt: now, UpdatedAt: now,
					Children: []model.Media{model.Media{Name: "plumb"}},
				},
				&model.Media{ID: 0, Name: "plumb", CreatedAt: now, UpdatedAt: now},
			},
			[]int{0},
			[][]int{[]int{0}},
			true,
		},
		{
			"inserts_parent_and_child_and_grandchild_successfully",
			model.Media{ID: 0, Name: "mango", Children: []model.Media{
				model.Media{ID: 0, Name: "plumb", Children: []model.Media{
					model.Media{ID: 0, Name: "grape"},
				}},
			}},
			[]*model.Media{nil, nil, nil},
			[]error{errors.New("blah"), errors.New("blah"), errors.New("blah")},
			nil,
			[]error{nil, nil, nil},
			[]error{nil, nil},
			nil,
			[]*model.Media{
				&model.Media{ID: 0, Name: "mango", CreatedAt: now, UpdatedAt: now,
					Children: []model.Media{model.Media{Name: "plumb",
						Children: []model.Media{model.Media{Name: "grape"}}},
					},
				},
				&model.Media{ID: 0, Name: "plumb", CreatedAt: now, UpdatedAt: now,
					Children: []model.Media{model.Media{Name: "grape"}},
				},
				&model.Media{ID: 0, Name: "grape", CreatedAt: now, UpdatedAt: now},
			},
			[]int{0, 0},
			[][]int{[]int{0}, []int{0}},
			true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := plogtest.SpannedTracedCtx()
			ml := plogtest.MockLogger{}

			ms := modeltest.MediaMockStore{
				GetByNameReturnMedia:                   tt.msGetByNameMedia,
				GetByNameReturnErr:                     tt.msGetByNameErr,
				UpdateReturnErr:                        tt.msUpdateErr,
				InsertReturnErr:                        tt.msInsertErr,
				AssociateParentIDWithChildIDsReturnErr: tt.msAssociateErr,
			}

			err := tt.m.Persist(ctx, &ml, &ms)
			if !tt.expErrToBeNil {
				if err == nil {
					t.Fatal("expected error to be non-nil, was nil")
				}

				return
			}

			if err != nil {
				t.Fatal("expected error to be nil, was: ", err)
			}

			if !cmp.Equal(tt.msUpdateExpMedia, ms.UpdateArgMedia) {
				t.Fatal("expected Update arg media not equal to actual: ",
					cmp.Diff(tt.msUpdateExpMedia, ms.UpdateArgMedia),
				)
			}
			if !cmp.Equal(tt.msInsertExpMedia, ms.InsertArgMedia) {
				t.Fatal("expected Insert arg media not equal to actual: ",
					cmp.Diff(tt.msInsertExpMedia, ms.InsertArgMedia),
				)
			}
			if !cmp.Equal(tt.msAssociateExpPID, ms.AssociateParentIDWithChildIDsArgPID) {
				t.Fatal("expected Associate arg pID not equal to actual: ",
					cmp.Diff(tt.msAssociateExpPID, ms.AssociateParentIDWithChildIDsArgPID),
				)
			}
			if !cmp.Equal(tt.msAssociateExpCID, ms.AssociateParentIDWithChildIDsArgCIDs) {
				t.Fatal("expected Associate arg cIDs not equal to actual: ",
					cmp.Diff(tt.msAssociateExpCID, ms.AssociateParentIDWithChildIDsArgCIDs),
				)
			}

		})
	}
}
