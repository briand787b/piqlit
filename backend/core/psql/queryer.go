package psql

import (
	"database/sql"

	"github.com/briand787b/piqlit/core/plog"
	"github.com/jmoiron/sqlx"
)

type queryLogger struct {
	logger  plog.Logger
	queryer sqlx.Queryer
}

func (ql *queryLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
	logQuery(ql.logger, query, args)
	return ql.queryer.Query(query, args...)
}

func (ql *queryLogger) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	logQuery(ql.logger, query, args)
	return ql.queryer.Queryx(query, args...)
}

func (ql *queryLogger) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	logQuery(ql.logger, query, args)
	return ql.queryer.QueryRowx(query, args...)
}
