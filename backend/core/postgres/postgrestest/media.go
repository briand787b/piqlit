package postgrestest

import (
	"fmt"

	"github.com/briand787b/piqlit/core/model"
	"github.com/briand787b/piqlit/core/obj"
	"github.com/briand787b/piqlit/core/plog/plogtest"
	"github.com/briand787b/piqlit/core/postgres"
)

// CreateMedia creates a Media record in the database
func (h *PGHelper) CreateMedia(m *model.Media, index int) *model.Media {
	if m == nil {
		m = &model.Media{
			Name:         fmt.Sprintf("%s_%v", h.T.Name(), index),
			Length:       1,
			Encoding:     obj.GIF,
			UploadStatus: obj.UploadInProgress,
			CreatedAt:    h.Tm,
			UpdatedAt:    h.Tm,
		}
	}

	ctx := plogtest.SpannedTracedCtx()
	ms := postgres.NewMediaPGStore(h.L, h.DB)
	if err := ms.Insert(
		ctx,
		m,
	); err != nil {
		defer h.Clean()
		h.T.Fatal("could not create Media: ", err)
	}

	h.L.Info(ctx, "created Media", "ID", m.ID)
	h.ParentIDs[Media] = m.ID
	return m
}
