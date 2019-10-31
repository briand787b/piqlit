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

	Content io.ReadCloser `db:"-" json:"-"`
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

// GetAllRootMedia x
func GetAllRootMedia(ctx context.Context, ms MediaStore) ([]Media, error) {
	mms, err := ms.GetAllRootMedia(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "could not get all root media from MediaStore")
	}

	return mms, nil
}

// Delete deletes the Media receiver from persistent storage
func (m *Media) Delete(ctx context.Context, l plog.Logger, mts MediaTxCtlStore) error {
	if m.ID == 0 {
		return perr.NewErrInvalid("cannot delete Media that is not persisted")
	}

	var err error
	if mts, err = mts.Begin(ctx); err != nil {
		return errors.Wrap(err, "could not begin tx")
	}

	if err := m.delete(ctx, mts); err != nil {
		if err := mts.Rollback(ctx); err != nil {
			l.Error(ctx, "could not roll back tx")
		}

		return errors.Wrap(err, "could not delete by ID")
	}

	if err := mts.Commit(ctx); err != nil {
		return errors.Wrap(err, "could not commit tx")
	}

	return nil
}

func (m *Media) delete(ctx context.Context, ms MediaStore) error {
	if err := ms.DisassociateParentIDFromChildren(ctx, m.ID); err != nil {
		return errors.Wrap(err, "could not disassociate parent from children")
	}

	if err := ms.DeleteByID(ctx, m.ID); err != nil {
		return errors.Wrap(err, "could not delete Media by ID")
	}

	return nil
}

// DownloadGZ returns a closable stream of the Media's contents
func (m *Media) DownloadGZ(ctx context.Context, l plog.Logger, os obj.ObjectStore) (io.ReadCloser, error) {
	if m.UploadStatus != obj.UploadDone {
		return nil, perr.NewErrNotFound(errors.New("requested Media has not completed uploading"))
	}

	rc, err := os.Get(ctx, m.Name)
	if err != nil {
		return nil, errors.Wrap(err, "could not get object from storage")
	}

	return rc, nil
}

// DownloadRaw returns the stored bytes in the native encoding they were uploaded in
func (m *Media) DownloadRaw(ctx context.Context, l plog.Logger, os obj.ObjectStore) (io.ReadCloser, error) {
	if m.UploadStatus != obj.UploadDone {
		return nil, perr.NewErrNotFound(errors.New("requested Media has not completed uploading"))
	}

	rc, err := os.Get(ctx, m.Name)
	if err != nil {
		return nil, errors.Wrap(err, "could not get object from storage")
	}

	dc, err := obj.NewGZDecompressor(ctx, l, rc)
	if err != nil {
		return nil, errors.Wrap(err, "could not create new gz decompressor")
	}

	return dc, nil
}

// Persist x
func (m *Media) Persist(ctx context.Context, l plog.Logger, mts MediaTxCtlStore) error {
	var err error
	mts, err = mts.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "could not begin tx")
	}

	if err := m.persist(ctx, l, mts); err != nil {
		if err := mts.Rollback(ctx); err != nil {
			l.Error(ctx, "could not roll back tx for Persist", "media", m, "error", err)
		}

		return errors.Wrap(err, "could not persist Media")
	}

	if err := mts.Commit(ctx); err != nil {
		l.Error(ctx, "could not commit tx", "error", err)
	}

	return nil
}

// Persist saves a Media to persistent storage
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

// UpdateShallow only updates the fields on the provided Media record, and not its children
func (m *Media) UpdateShallow(ctx context.Context, l plog.Logger, ms MediaStore) error {
	if err := m.Validate(ctx, l); err != nil {
		return errors.Wrap(err, "could not validate Media")
	}

	return m.update(ctx, l, ms)
}

// UploadRaw uploads the provided contents to object storage.  Objects are stored in a gzipped format
func (m *Media) UploadRaw(ctx context.Context, l plog.Logger, ms MediaStore, os obj.ObjectStore) error {
	if m.Length < 1 {
		return perr.NewErrInvalid("zero-length files cannot be uploaded to - update record's length first")
	}

	m.UploadStatus = obj.UploadInProgress
	if err := ms.Update(ctx, m); err != nil {
		return errors.Wrap(err, "could not update upload_status")
	}

	gc := obj.NewGZCompressor(ctx, l, m.Content)
	defer l.Close(ctx, gc)

	if err := os.Put(ctx, m.Name, gc); err != nil {
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
func (m *Media) insert(ctx context.Context, l plog.Logger, ms MediaStore) error {
	if _, err := ms.GetByName(ctx, m.Name); err == nil {
		return perr.NewErrInvalid(fmt.Sprintf("name '%s' already exists in database", m.Name))
	}

	if m.Length == 0 {
		m.UploadStatus = obj.UploadDone
	} else {
		m.UploadStatus = obj.UploadNotStarted
	}

	now := time.Now().UTC().Truncate(time.Second)
	m.CreatedAt, m.UpdatedAt = now, now
	return ms.Insert(ctx, m)
}

// only update media with duplicate name if same ID
func (m *Media) update(ctx context.Context, l plog.Logger, ms MediaStore) error {
	mm, err := ms.GetByID(ctx, m.ID)
	if err != nil {
		return perr.NewErrNotFound(errors.Wrap(err, "could not find "))
	}

	// do not update media that are mid-upload
	if mm.UploadStatus == obj.UploadInProgress {
		return perr.NewErrInvalid("cannot update media that is currently being uploaded")
	}

	l.Info(ctx, "Media to update", "pre-update value", mm)

	if mm.Name != m.Name {
		if _, err = ms.GetByName(ctx, m.Name); err == nil {
			return perr.NewErrInvalid(fmt.Sprintf("name '%s' already exists in database", m.Name))
		}
	}

	if m.Length == 0 {
		// if obj has no content, it should always have an upload status of 'done'
		m.UploadStatus = obj.UploadDone
	} else if mm.Length < 1 {
		// if obj had no content, but now does, its previously 'done' upload status is now invalid
		m.UploadStatus = obj.UploadNotStarted
	} else if m.Length != mm.Length {
		// if content length changed, previous upload status is now invalid
		m.UploadStatus = obj.UploadNotStarted
	} else {
		// in all other cases, the updated object has the same status it previously had
		m.UploadStatus = mm.UploadStatus
	}

	m.CreatedAt = mm.CreatedAt
	m.UpdatedAt = time.Now().UTC().Truncate(time.Second)
	return ms.Update(ctx, m)
}
