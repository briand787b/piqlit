package postgres

import (
	"context"
	"fmt"

	"github.com/briand787b/piqlit/core/model"
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

// DisassociateParentIDFromChildIDs deletes records from the parent_child_media table where the parent_id
// equals pID and the cID is in the set of cIDs
func (mps *MediaPGStore) DisassociateParentIDFromChildIDs(ctx context.Context, pID int, cIDs ...int) error {
	if cIDs == nil || len(cIDs) < 1 {
		return nil
	}

	qry, args, err := sqlx.In(`
		DELETE FROM parent_child_media 
		WHERE 
			parent_id = ?
			AND child_id IN (?)
		RETURNING child_id;`,
		pID, cIDs,
	)
	if err != nil {
		return errors.Wrap(err, "could not format `IN` query")
	}

	qry = mps.db.Rebind(qry)

	var deletedChildIDs []int
	if err := sqlx.SelectContext(ctx, mps.db, &deletedChildIDs, qry, args...); err != nil {
		return errors.Wrap(err, "could not execute query")
	}

	return nil
}

// SelectByParentID returns a slice of model.Media with the provided parentID
func (mps *MediaPGStore) SelectByParentID(ctx context.Context, pID int) ([]model.Media, error) {
	var m []model.Media
	if err := sqlx.SelectContext(ctx, mps.db, &m, `
		SELECT
			m.id,
			m.name,
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

	return m, nil
}

// GetByID returns a Media record by its id
func (mps *MediaPGStore) GetByID(ctx context.Context, id int) (*model.Media, error) {
	var m model.Media
	if err := sqlx.GetContext(ctx, mps.db, &m, `
	SELECT
		id,
		name,
		encoding,
		upload_status,
		created_at,
		updated_at
	FROM media
	WHERE id = $1;`,
		id,
	); err != nil {
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
			encoding,
			upload_status,
			created_at,
			updated_at
		)
		VALUES
		(
			:name,
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
