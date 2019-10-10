package model

import (
	"context"
	"time"

	"github.com/briand787b/piqlit/core/obj"
	"github.com/briand787b/piqlit/core/perr"
	"github.com/briand787b/piqlit/core/plog"

	"github.com/pkg/errors"
)

// Media is a container for viewable works of art, or for child media
type Media struct {
	ID           int              `db:"id"`
	Name         string           `db:"name"`
	Length       int64            `db:"length"`
	Encoding     obj.Encoding     `db:"encoding"`
	UploadStatus obj.UploadStatus `db:"upload_status"`
	Children     []Media          `db:"children"`
	CreatedAt    time.Time        `db:"created_at"`
	UpdatedAt    time.Time        `db:"updated_at"`
}

// Persist saves a Media to persistent storage
func (m *Media) Persist(ctx context.Context, l plog.Logger, ms MediaStore) error {
	if err := m.Validate(ctx, l); err != nil {
		return errors.Wrap(err, "could not validate Media")
	}

	if m.ID > 0 {
		if err := m.update(ctx, l, ms); err != nil {
			return errors.Wrap(err, "could not update Media")
		}
	} else {
		if err := m.insert(ctx, l, ms); err != nil {
			return errors.Wrap(err, "could not insert Media")
		}
	}

	var cIDs []int
	for _, c := range m.Children {
		if err := c.Persist(ctx, l, ms); err != nil {
			return errors.Wrap(err, "could not persist children")
		}

		cIDs = append(cIDs, c.ID)
	}

	if err := ms.AssociateParentIDWithChildIDs(ctx, m.ID, cIDs...); err != nil {
		return errors.Wrap(err, "could not associate parent media with children")
	}

	return nil
}

// Validate returns an error if the Media is not properly
// configured for persistent storage
func (m *Media) Validate(ctx context.Context, l plog.Logger) error {
	if m.Name == "" {
		l.Invalid(ctx, *m, "empty field: name")
		return perr.NewErrInvalid("Media.Name cannot be empty")
	}

	if m.Length < 1 {
		l.Invalid(ctx, *m, "non-positive length")
		return perr.NewErrInvalid("Media.Length must be positive integer")
	}

	return nil
}

func (m *Media) update(ctx context.Context, l plog.Logger, ms MediaStore) error {
	m.UpdatedAt = time.Now()
	return ms.Update(ctx, m)
}

func (m *Media) insert(ctx context.Context, l plog.Logger, ms MediaStore) error {
	now := time.Now()
	m.CreatedAt, m.UpdatedAt = now, now
	return ms.Insert(ctx, m)
}
