package psql

import (
	"database/sql"

	"github.com/briand787b/piqlit/core/plog"

	"github.com/jmoiron/sqlx"
)

type execLogger struct {
	logger plog.Logger
	execer sqlx.Execer
}

func (el *execLogger) Exec(query string, args ...interface{}) (sql.Result, error) {
	logQuery(el.logger, query, args)
	return el.execer.Exec(query, args...)
}
