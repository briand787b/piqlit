package model

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/briand787b/piqlit/core/obj"
	"github.com/briand787b/piqlit/core/perr"
	"github.com/briand787b/piqlit/core/plog"

	"github.com/pkg/errors"
)

// Media is a container for viewable works of art, or for child media
type Media struct {
	ID           int              `db:"id" json:"id"`
	Name         string           `db:"name" json:"name"`
	Length       int64            `db:"length" json:"length"`
	Encoding     obj.Encoding     `db:"encoding" json:"encoding"`
	UploadStatus obj.UploadStatus `db:"upload_status" json:"upload_status"`
	Children     []Media          `db:"children" json:"children"`
	CreatedAt    time.Time        `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time        `db:"updated_at" json:"updated_at"`
}

// FindMediaByID returns the Media with provided id
func FindMediaByID(ctx context.Context, ms MediaStore, id int) (*Media, error) {
	if id == 0 {
		return nil, perr.NewErrInvalid("cannot search for ID 0")
	}

	m, err := ms.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "could not get Media by ID")
	}

	cs, err := ms.SelectByParentID(ctx, m.ID)
	if err != nil {
		return nil, errors.Wrap(err, "could not get child Media")
	}

	var cc *Media
	for i, c := range cs {
		cc, err = FindMediaByID(ctx, ms, c.ID)
		if err != nil {
			return nil, errors.Wrap(err, "could not get grandchild media")
		}

		cs[i].Children = cc.Children
	}

	m.Children = cs
	return m, nil
}

// Delete deletes the Media receiver from persistent storage
func (m *Media) Delete(ctx context.Context, ms MediaStore) error {
	if m.ID == 0 {
		return perr.NewErrInvalid("cannot delete Media that is not persisted")
	}

	if err := ms.DeleteByID(ctx, m.ID); err != nil {
		return errors.Wrap(err, "could not delete Media by ID")
	}

	return nil
}

// Download returns a closable stream of the Media's contents
func (m *Media) Download(ctx context.Context, l plog.Logger, os obj.ObjectStore) (io.ReadCloser, error) {
	if m.UploadStatus != obj.UploadDone {
		return nil, perr.NewErrNotFound(errors.New("requested Media has not completed uploading"))
	}

	rc, err := os.Get(ctx, m.Name)
	if err != nil {
		return nil, errors.Wrap(err, "could not get object from storage")
	}

	return rc, nil
}

// Persist x
func (m *Media) Persist(ctx context.Context, l plog.Logger, mts MediaTxCtlStore) error {
	txCtl, err := mts.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "could not begin tx")
	}

	if err := m.persist(ctx, l, mts); err != nil {
		if err := txCtl.Rollback(ctx); err != nil {
			l.Error(ctx, "could not roll back tx for Persist", "media", m, "error", err)
		}

		return errors.Wrap(err, "could not persist Media")
	}

	if err := txCtl.Commit(ctx); err != nil {
		l.Error(ctx, "could not commit tx", "error", err)
	}

	return nil
}

// Persist saves a Media to persistent storage
//
// TODO: figure out rollback strategy
func (m *Media) persist(ctx context.Context, l plog.Logger, ms MediaStore) error {
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

	if err := ms.DisassociateParentIDFromChildren(ctx, m.ID); err != nil {
		return errors.Wrap(err, "could not disassociate parent Media from children")
	}

	var cIDs []int
	for _, c := range m.Children {
		if err := c.persist(ctx, l, ms); err != nil {
			return errors.Wrap(err, "could not persist children")
		}

		cIDs = append(cIDs, c.ID)
	}

	if m.Children != nil && len(m.Children) > 0 {
		if err := ms.AssociateParentIDWithChildIDs(ctx, m.ID, cIDs...); err != nil {
			return errors.Wrap(err, "could not associate parent media with children")
		}
	}

	return nil
}

// Upload uploads the provided contents to object storage
func (m *Media) Upload(ctx context.Context, l plog.Logger, os obj.ObjectStore, ms MediaStore, content io.ReadCloser) error {
	if m.Length < 1 {
		// zero-length objects cannot be uploaded - unsure if this is an error or not though
		return nil
	}

	m.UploadStatus = obj.UploadInProgress
	if err := ms.Update(ctx, m); err != nil {
		return errors.Wrap(err, "could not update upload_status")
	}

	if err := os.Put(ctx, m.Name, content); err != nil {
		m.UploadStatus = obj.UploadFailed
		if err := ms.Update(ctx, m); err != nil {
			l.Error(ctx, "could not update media fields", "UploadStatus", obj.UploadFailed)
		}

		return errors.Wrap(err, "could not store object")
	}

	m.UploadStatus = obj.UploadDone
	if err := ms.Update(ctx, m); err != nil {
		return errors.Wrap(err, "could not update upload_status")
	}

	return nil
}

// Validate returns an error if the Media is not properly
// configured for persistent storage
//
// zero-length media may exist as org units
func (m *Media) Validate(ctx context.Context, l plog.Logger) error {
	if m.Name == "" {
		l.Invalid(ctx, *m, "empty field: name")
		return perr.NewErrInvalid("Media.Name cannot be empty")
	}

	if m.Length < 0 {
		l.Invalid(ctx, *m, "non-positive length")
		return perr.NewErrInvalid("Media.Length must be positive integer")
	}

	return nil
}

// do not insert media with existing names
//
// TODO: set upload_status to `not_started` before insertion
func (m *Media) insert(ctx context.Context, l plog.Logger, ms MediaStore) error {
	if _, err := ms.GetByName(ctx, m.Name); err == nil {
		return perr.NewErrInvalid(fmt.Sprintf("name '%s' already exists in database", m.Name))
	}

	now := time.Now().UTC().Truncate(time.Second)
	m.CreatedAt, m.UpdatedAt = now, now
	m.UploadStatus = obj.UploadNotStarted
	return ms.Insert(ctx, m)
}

// only update media with duplicate name if same ID
func (m *Media) update(ctx context.Context, l plog.Logger, ms MediaStore) error {
	mm, err := ms.GetByID(ctx, m.ID)
	if err != nil {
		return perr.NewErrNotFound(errors.Wrap(err, "could not find "))
	}

	l.Info(ctx, "Media to update", "pre-update value", mm)

	if mm.Name != m.Name {
		if _, err = ms.GetByName(ctx, m.Name); err == nil {
			return perr.NewErrInvalid(fmt.Sprintf("name '%s' already exists in database", m.Name))
		}
	}

	m.UploadStatus = mm.UploadStatus
	m.CreatedAt = mm.CreatedAt
	m.UpdatedAt = time.Now().UTC().Truncate(time.Second)
	return ms.Update(ctx, m)
}
