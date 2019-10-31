package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/briand787b/piqlit/core/model"
	"github.com/briand787b/piqlit/core/perr"
	"github.com/briand787b/piqlit/core/plog"
	"github.com/briand787b/piqlit/core/psql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// MediaPGStore is a MediaStore backed by Po
type MediaPGStore struct {
	l  plog.Logger
	db psql.ExtFull
}

// NewMediaPGStore is a MediaStore backed by Postgres
func NewMediaPGStore(l plog.Logger, db psql.ExtFull) *MediaPGStore {
	return &MediaPGStore{l: l, db: db}
}

// Begin x
func (mps *MediaPGStore) Begin(ctx context.Context) (model.MediaTxCtlStore, error) {
	tx, err := mps.db.Begin(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not begin txx")
	}

	return NewMediaPGStore(mps.l, psql.NewExtFullFromTx(mps.l, tx)), nil
}

// Commit x
func (mps *MediaPGStore) Commit(ctx context.Context) error {
	return mps.db.Commit(ctx)
}

// Rollback x
func (mps *MediaPGStore) Rollback(ctx context.Context) error {
	return mps.db.Rollback(ctx)
}

// AssociateParentIDWithChildIDs inserts a row into the parent_child_media table
func (mps *MediaPGStore) AssociateParentIDWithChildIDs(ctx context.Context, pID int, cIDs ...int) error {
	if cIDs == nil || len(cIDs) < 1 {
		return nil
	}

	baseQ := `INSERT INTO parent_child_media
	(
		parent_id,
		child_id
	)
	VALUES
	`

	args := make([]interface{}, 2*len(cIDs))
	var cnt int
	for i, cID := range cIDs {
		cnt = i * 2
		baseQ += fmt.Sprintf("($%v, $%v),\n", cnt+1, cnt+2)
		args[cnt] = pID
		args[cnt+1] = cID
	}

	baseQ = baseQ[:len(baseQ)-2] + " RETURNING parent_id;"

	var id int
	if err := sqlx.GetContext(ctx, mps.db, &id, baseQ, args...); err != nil {
		return errors.Wrap(err, "could not execute query")
	}

	return nil
}

// DeleteByID deletes a Media record by its id
func (mps *MediaPGStore) DeleteByID(ctx context.Context, id int) error {
	var i int
	if err := sqlx.GetContext(ctx, mps.db, &i, `
		DELETE FROM media
		WHERE
			id = $1
		RETURNING id;`,
		id,
	); err != nil {
		return errors.Wrap(err, "could not execute query")
	}

	return nil
}

// DisassociateParentIDFromChildren deletes records from the parent_child_media table where the parent_id
// equals pID
func (mps *MediaPGStore) DisassociateParentIDFromChildren(ctx context.Context, pID int) error {
	var CIDs []int64
	if err := sqlx.SelectContext(ctx, mps.db, &CIDs, `
		DELETE FROM parent_child_media
		WHERE parent_id = $1
		RETURNING child_id;`,
		pID,
	); err != nil {
		return errors.Wrap(err, "could not execute query")
	}

	return nil
}

// GetAllRootMedia returns all Media that are not children of other Media
func (mps *MediaPGStore) GetAllRootMedia(ctx context.Context) ([]model.Media, error) {
	var ms []model.Media
	if err := sqlx.SelectContext(ctx, mps.db, &ms, `
		SELECT
			id,
			name,
			length,
			encoding,
			upload_status,
			created_at,
			updated_at
		FROM
			media m
		LEFT OUTER JOIN parent_child_media pcm
		ON m.id = pcm.child_id
		WHERE
			pcm.child_id IS NULL;`,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, perr.NewErrNotFound(err)
		}

		return nil, errors.Wrap(err, "could not execute query")
	}

	return ms, nil
}

// GetByID returns a Media record by its id
func (mps *MediaPGStore) GetByID(ctx context.Context, id int) (*model.Media, error) {
	var m model.Media
	if err := sqlx.GetContext(ctx, mps.db, &m, `
	SELECT
		id,
		name,
		length,
		encoding,
		upload_status,
		created_at,
		updated_at
	FROM media
	WHERE id = $1;`,
		id,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, perr.NewErrNotFound(err)
		}

		return nil, perr.NewErrInternal(err)
	}

	return &m, nil
}

// GetByName retrieves a Media by its name
func (mps *MediaPGStore) GetByName(ctx context.Context, name string) (*model.Media, error) {
	var m model.Media
	if err := sqlx.GetContext(ctx, mps.db, &m, `
		SELECT
			id,
			name,
			length,
			encoding,
			upload_status,
			created_at,
			updated_at
		FROM media
		WHERE name = $1;`,
		name,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, perr.NewErrNotFound(err)
		}

		return nil, errors.Wrap(err, "could not execute query")
	}

	return &m, nil
}

// Insert inserts the Media record into Postgres
func (mps *MediaPGStore) Insert(ctx context.Context, m *model.Media) error {
	qry, args, err := sqlx.Named(`
		INSERT INTO	media
		(
			name,
			length,
			encoding,
			upload_status,
			created_at,
			updated_at
		)
		VALUES
		(
			:name,
			:length,
			:encoding,
			:upload_status,
			:created_at,
			:updated_at
		)
		RETURNING id;`,
		*m,
	)
	if err != nil {
		return errors.Wrap(err, "could not bind Media to query")
	}

	qry = sqlx.Rebind(sqlx.DOLLAR, qry)

	var saveID int
	if err := sqlx.GetContext(ctx, mps.db, &saveID, qry, args...); err != nil {
		return errors.Wrap(err, "could not execute query")
	}

	m.ID = saveID
	return nil
}

// SelectByParentID returns a slice of model.Media with the provided parentID
func (mps *MediaPGStore) SelectByParentID(ctx context.Context, pID int) ([]model.Media, error) {
	var ms []model.Media
	if err := sqlx.SelectContext(ctx, mps.db, &ms, `
		SELECT
			m.id,
			m.name,
			m.length,
			m.encoding,
			m.upload_status,
			m.created_at,
			m.updated_at
		FROM media m
		INNER JOIN parent_child_media pcm
		ON m.id = pcm.child_id
		WHERE pcm.parent_id = $1;`,
		pID,
	); err != nil {
		return nil, errors.Wrap(err, "could not execute query")
	}

	return ms, nil
}

// Update updates a model.Media record
func (mps *MediaPGStore) Update(ctx context.Context, m *model.Media) error {
	qry, args, err := sqlx.Named(`
		UPDATE media
		SET
			name = :name,
			length = :length,
			encoding = :encoding,
			upload_status = :upload_status,
			updated_at = :updated_at
		WHERE
			id = :id
		RETURNING id;`, *m)
	if err != nil {
		return errors.Wrap(err, "could not bind Tag to query")
	}

	qry = sqlx.Rebind(sqlx.DOLLAR, qry)

	var updateID int
	if err := sqlx.GetContext(ctx, mps.db, &updateID, qry, args...); err != nil {
		return errors.Wrap(err, "could not execute query")
	}

	return nil
}
