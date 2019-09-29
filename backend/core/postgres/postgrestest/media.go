package postgrestest

import (
	"context"
	"fmt"

	"github.com/briand787b/piqlit/core/model"
	"github.com/briand787b/piqlit/core/obj"
	"github.com/briand787b/piqlit/core/postgres"
)

// CreateMedia creates a Media record in the database
func (h *PGHelper) CreateMedia(m *model.Media, index int) *model.Media {
	if m == nil {
		m = &model.Media{
			Name:         fmt.Sprintf("%s_%v", h.T.Name(), index),
			Encoding:     obj.GIF,
			UploadStatus: obj.UploadInProgress,
			CreatedAt:    h.Tm,
			UpdatedAt:    h.Tm,
		}
	}

	ms := postgres.NewMediaPGStore(h.L, h.DB)
	if err := ms.Insert(
		context.Background(),
		m,
	); err != nil {
		defer h.Clean()
		h.T.Fatal("could not create Media: ", err)
	}

	h.L.Infow("created Media", "ID", m.ID)

	h.CF.Add(func() {
		if err := ms.DeleteByID(context.Background(), m.ID); err != nil {
			h.T.Fatal("could not delete Media")
		}
	})

	h.ParentIDs[Media] = m.ID
	return m
}