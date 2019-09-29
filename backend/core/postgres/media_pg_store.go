package postgres

import (
	"context"

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

// DeleteByID deletes a Media by its id
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
		FROM 
			media
		WHERE
			id = $1;`,
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
