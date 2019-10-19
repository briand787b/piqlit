package psql

import (
	"context"
	"database/sql"

	"github.com/briand787b/piqlit/core/plog"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type txFull interface {
	Begin(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type txBeginner struct {
	l  plog.Logger
	db sqlx.DB
}

func (t *txBeginner) Begin(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	t.l.Info(ctx, "beginning transaction")
	return t.db.BeginTxx(ctx, opts)
}

// This is effectively a nop to satisfy the interface
func (t *txBeginner) Commit(ctx context.Context) error {
	return errors.New("attempting to Commit non-Tx DB connection")
}

// This is effectively a nop to satisfy the interface
func (t *txBeginner) Rollback(ctx context.Context) error {
	return errors.New("attempting to Rollback non-Tx DB connection")
}

type txCloser struct {
	l plog.Logger
	t sqlx.Tx
}

// This is effectively a nop to satisfy the interface
func (t *txCloser) Begin(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	return nil, errors.New("attempting to begin an already begun Tx")
}

func (t *txCloser) Commit(ctx context.Context) error {
	t.l.Info(ctx, "committing tx")
	if err := t.t.Commit(); err != nil {
		if err := t.Rollback(ctx); err != nil {
			t.l.Error(ctx, "could not roll back uncommittable tx")
		}

		return errors.Wrap(err, "could not commit tx")
	}

	return nil
}

func (t *txCloser) Rollback(ctx context.Context) error {
	t.l.Info(ctx, "rolling tx back")
	if err := t.t.Rollback(); err != nil {
		return errors.New("could not roll back tx")
	}

	return nil
}
