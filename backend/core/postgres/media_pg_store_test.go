package postgres_test

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/briand787b/piqlit/core/model"
	"github.com/briand787b/piqlit/core/obj"
	"github.com/briand787b/piqlit/core/perr"
	"github.com/briand787b/piqlit/core/plog/plogtest"
	"github.com/briand787b/piqlit/core/postgres"
	"github.com/briand787b/piqlit/core/postgres/postgrestest"
	"github.com/briand787b/piqlit/core/test"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pkg/errors"
)

func TestMediaPGStoreGetByID(t *testing.T) {
	test.SkipLong(t)
	tests := []struct {
		name       string
		count      int
		idValid    bool
		expBaseErr error
		expM       model.Media
	}{
		{
			"finding_valid_id_in_single_successful",
			1,
			true,
			nil,
			model.Media{
				Name:         "TestMediaPGStoreGetByID/finding_valid_id_in_single_successful_0",
				Length:       1,
				Encoding:     obj.GIF,
				UploadStatus: obj.UploadInProgress,
			},
		},
		{
			"finding_valid_id_in_multi_successful",
			3,
			true,
			nil,
			model.Media{
				Name:         "TestMediaPGStoreGetByID/finding_valid_id_in_multi_successful_2",
				Length:       1,
				Encoding:     obj.GIF,
				UploadStatus: obj.UploadInProgress,
			},
		},
		{
			"finding_invalid_id_in_single_fails",
			1,
			false,
			perr.ErrNotFound,
			model.Media{},
		},
		{
			"finding_invalid_id_in_multi_fails",
			3,
			false,
			perr.ErrNotFound,
			model.Media{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := postgrestest.NewPGHelper(t)
			var id int
			for i := 0; i < tt.count; i++ {
				if tt.idValid {
					id = h.CreateMedia(nil, i).ID
				}
			}

			defer h.Clean()

			retM, err := postgres.NewMediaPGStore(h.L, h.DB).GetByID(
				context.Background(),
				id,
			)

			if e := errors.Cause(err); !cmp.Equal(tt.expBaseErr, e) {
				t.Fatal("expected error not equal to returned error: ",
					cmp.Diff(tt.expBaseErr, e),
				)
			}

			if tt.expBaseErr != nil {
				return
			}

			tt.expM.ID = id
			tt.expM.CreatedAt, tt.expM.UpdatedAt = h.Tm, h.Tm
			if !cmp.Equal(tt.expM, *retM) {
				t.Fatalf("expected Media and returned media are not equal: %s",
					cmp.Diff(tt.expM, *retM),
				)
			}
		})
	}
}

func TestMediaPGStoreInsertDelete(t *testing.T) {
	test.SkipLong(t)
	tests := []struct {
		name string
		m    model.Media
	}{
		{
			"successful_with_all_fields_filled_in",
			model.Media{
				Name:         "mango",
				Encoding:     obj.GIF,
				Length:       1,
				UploadStatus: obj.UploadInProgress,
				CreatedAt:    time.Now().UTC().Truncate(time.Second),
				UpdatedAt:    time.Now().UTC().Truncate(time.Second),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := postgrestest.NewPGHelper(t)
			mps := postgres.NewMediaPGStore(h.L, h.DB)
			ctx := context.Background()

			if err := mps.Insert(ctx, &tt.m); err != nil {
				t.Fatal(err)
			}

			retM, err := mps.GetByID(ctx, tt.m.ID)
			if err != nil {
				t.Fatal("could not get by id: ", err)
			}

			tt.m.ID = retM.ID
			if !cmp.Equal(tt.m, *retM) {
				t.Fatalf("expected Media not equal to returned media %s",
					cmp.Diff(tt.m, *retM),
				)
			}

			if err := mps.DeleteByID(ctx, retM.ID); err != nil {
				t.Fatal(err)
			}

			if _, err := mps.GetByID(ctx, tt.m.ID); err == nil {
				t.Fatal("expected error to be non-nil, was nil")
			}
		})
	}
}

func TestMediaParentChildAssociateDisassociateByID(t *testing.T) {
	test.SkipLong(t)
	tests := []struct {
		name        string
		numChildren int
	}{
		{
			"associates_to_one_child",
			1,
		},
		{
			"associates_to_two_children",
			2,
		},
		{
			"associates_to_three_children",
			3,
		},
		{
			"associates_to_no_children",
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := postgrestest.NewPGHelper(t)
			pm := h.CreateMedia(nil, 0)
			defer h.Clean()
			cms := make([]model.Media, tt.numChildren)
			cmIDs := make([]int, tt.numChildren)
			for i := 0; i < tt.numChildren; i++ {
				cms[i] = *h.CreateMedia(nil, i+1)
				cmIDs[i] = cms[i].ID
			}

			mps := postgres.NewMediaPGStore(h.L, h.DB)
			ctx := context.Background()

			if err := mps.AssociateParentIDWithChildIDs(ctx, pm.ID, cmIDs...); err != nil {
				t.Fatal("failed to prove parent-child media association: ", err)
			}

			defer mps.DisassociateParentIDFromChildIDs(ctx, pm.ID, cmIDs...)

			retCMs, err := mps.SelectByParentID(ctx, pm.ID)
			if err != nil {
				t.Fatal("failed to get associated parent media with children: ", err)
			}

			if !cmp.Equal(cms, retCMs, cmpopts.EquateEmpty()) {
				t.Fatal("expected ChildMedia not equal to returned child media: ",
					cmp.Diff(cms, retCMs, cmpopts.EquateEmpty()),
				)
			}

			if err := mps.DisassociateParentIDFromChildIDs(ctx, pm.ID, cmIDs...); err != nil {
				t.Fatal("could not disassociate parent media from children: ", err)
			}

			retCMs, err = mps.SelectByParentID(ctx, pm.ID)
			if err != nil {
				t.Fatal("failed to prove parent-child media disassociation: ", err)
			}

			if len(retCMs) != 0 {
				t.Fatalf("expected returned child media to equal 0, returned %v", len(retCMs))
			}
		})
	}
}

func TestEncodings(t *testing.T) {
	test.SkipLong(t)
	tests := []struct {
		name        string
		encodingStr string
		expEncoding obj.Encoding
	}{
		{"gif_encoding", "gif", obj.GIF},
		{"empty_encoding", "", obj.Empty},
		{"unknown_encoding", "unknown", obj.Encoding("unknown")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := postgrestest.NewPGHelper(t)
			m := model.Media{
				Name:         t.Name(),
				Length:       1,
				Encoding:     obj.Encoding(tt.encodingStr),
				UploadStatus: obj.UploadInProgress,
				CreatedAt:    h.Tm,
				UpdatedAt:    h.Tm,
			}

			h.CreateMedia(&m, 0)
			defer h.Clean()

			retM, err := postgres.NewMediaPGStore(h.L, h.DB).GetByID(
				plogtest.SpannedTracedCtx(),
				m.ID,
			)

			if err != nil {
				t.Fatal(err)
			}

			if retM.Encoding != tt.expEncoding {
				t.Fatalf("expected encoding to be %v, was %v",
					retM.Encoding,
					tt.expEncoding,
				)
			}
		})
	}
}

func TestMediaPGStoreUpdate(t *testing.T) {
	test.SkipLong(t)
	now := time.Now().UTC().Truncate(time.Second)
	tests := []struct {
		name     string
		updMedia model.Media
	}{
		{
			"empty_struct",
			model.Media{
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			"full_name",
			model.Media{
				Name:      "name",
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			"full_encoding",
			model.Media{
				Encoding:  obj.GIF,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			"full_upload_status",
			model.Media{
				UploadStatus: obj.UploadDone,
				CreatedAt:    now,
				UpdatedAt:    now,
			},
		},
		{
			"full_length",
			model.Media{
				Length:    math.MaxInt64,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := postgrestest.NewPGHelper(t)
			tt.updMedia.ID = h.CreateMedia(nil, 0).ID
			defer h.Clean()

			ctx := plogtest.SpannedTracedCtx()
			mps := postgres.NewMediaPGStore(h.L, h.DB)

			if err := mps.Update(ctx, &tt.updMedia); err != nil {
				t.Fatal(err)
			}

			retM, err := mps.GetByID(ctx, tt.updMedia.ID)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(tt.updMedia, *retM) {
				t.Fatal("expected and retrieved media different: ",
					cmp.Diff(tt.updMedia, *retM),
				)
			}
		})
	}
}
