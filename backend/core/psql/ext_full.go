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
}

// GetExtFull returns an implementation of ExtFull backed by
// whatever database underlies the db variable provided to it.
func GetExtFull(l *plog.Logger, db *sqlx.DB) ExtFull {
	return struct {
		binder
		sqlx.Execer
		sqlx.ExecerContext
		sqlx.Queryer
		sqlx.QueryerContext
	}{
		db,
		&execLogger{
			logger: *l,
			execer: db,
		},
		&execContextLogger{
			logger:        *l,
			execerContext: db,
		},
		&queryLogger{
			logger:  *l,
			queryer: db,
		},
		&queryContextLogger{
			logger:         *l,
			queryerContext: db,
		},
	}
}
