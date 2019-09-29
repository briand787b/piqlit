package postgres_test

import (
	"context"
	"testing"

	"github.com/briand787b/piqlit/core/model"
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
			model.Media{},
		},
		{
			"finding_valid_id_in_multi_successful",
			3,
			true,
			true,
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
