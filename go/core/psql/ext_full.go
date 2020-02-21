package psql

import (
	"github.com/briand787b/piqlit/core/plog"

	"github.com/jmoiron/sqlx"
)

// ExtFull is the interface that abstracts all the querying
// and executing functions of the sqlx package.  It is satisfied
// by the sqlx.DB type.
type ExtFull interface {
	binder
	sqlx.Execer
	sqlx.ExecerContext
	sqlx.Queryer
	sqlx.QueryerContext
	txFull
}

// NewExtFull returns an implementation of ExtFull backed by
// whatever database underlies the db variable provided to it.
func NewExtFull(l plog.Logger, db *sqlx.DB) ExtFull {
	return struct {
		binder
		sqlx.Execer
		sqlx.ExecerContext
		sqlx.Queryer
		sqlx.QueryerContext
		txFull
	}{
		db,
		&execLogger{
			logger: l,
			execer: db,
		},
		&execContextLogger{
			logger:        l,
			execerContext: db,
		},
		&queryLogger{
			logger:  l,
			queryer: db,
		},
		&queryContextLogger{
			logger:         l,
			queryerContext: db,
		},
		&txBeginner{
			l:  l,
			db: *db,
		},
	}
}

// NewExtFullFromTx returns an implementation of ExtFullTx backed by
// whatever database underlies the db variable provided to it.
func NewExtFullFromTx(l plog.Logger, tx *sqlx.Tx) ExtFull {
	return struct {
		binder
		sqlx.Execer
		sqlx.ExecerContext
		sqlx.Queryer
		sqlx.QueryerContext
		txFull
	}{
		tx,
		&execLogger{
			logger: l,
			execer: tx,
		},
		&execContextLogger{
			logger:        l,
			execerContext: tx,
		},
		&queryLogger{
			logger:  l,
			queryer: tx,
		},
		&queryContextLogger{
			logger:         l,
			queryerContext: tx,
		},
		&txCloser{
			l: l,
			t: *tx,
		},
	}
}
