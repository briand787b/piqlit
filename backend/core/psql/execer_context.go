package psql

import (
	"context"
	"database/sql"

	"github.com/briand787b/piqlit/core/plog"
	"github.com/jmoiron/sqlx"
)

type execContextLogger struct {
	logger        plog.Logger
	execerContext sqlx.ExecerContext
}

func (ecl *execContextLogger) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	logQuery(ecl.logger, query, args)
	return ecl.execerContext.ExecContext(ctx, query, args)
}
