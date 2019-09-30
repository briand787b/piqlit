package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/briand787b/piqlit/core/model"
	"github.com/briand787b/piqlit/core/obj"
	"github.com/briand787b/piqlit/core/postgres"
	"github.com/briand787b/piqlit/core/postgres/postgrestest"

	"github.com/google/go-cmp/cmp"
)

func TestMediaPGStoreGetByID(t *testing.T) {
	tests := []struct {
		name          string
		count         int
		idValid       bool
		expErrToBeNil bool
		expM          model.Media
	}{
		{
			"finding_valid_id_in_single_successful",
			1,
			true,
			true,
			model.Media{
				Name:         "TestMediaPGStoreGetByID/finding_valid_id_in_single_successful_0",
				Encoding:     obj.GIF,
				UploadStatus: obj.UploadInProgress,
			},
		},
		{
			"finding_valid_id_in_multi_successful",
			3,
			true,
			true,
			model.Media{
				Name:         "TestMediaPGStoreGetByID/finding_valid_id_in_multi_successful_2",
				Encoding:     obj.GIF,
				UploadStatus: obj.UploadInProgress,
			},
		},
		{
			"finding_invalid_id_in_single_fails",
			1,
			false,
			false,
			model.Media{},
		},
		{
			"finding_invalid_id_in_multi_fails",
			3,
			false,
			false,
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

			if !tt.expErrToBeNil {
				if err == nil {
					t.Fatal("expected error to be non-nil, was nil")
				}

				return
			}

			if err != nil {
				t.Fatal("expected error to be nil, was ", err)
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
	tests := []struct {
		name string
		m    model.Media
	}{
		{
			"successful_with_all_fields_filled_in",
			model.Media{
				Name:         "mango",
				Encoding:     obj.GIF,
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
