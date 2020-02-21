package psql

import (
	"context"
	"database/sql"

	"github.com/briand787b/piqlit/core/plog"
	"github.com/jmoiron/sqlx"
)

type queryContextLogger struct {
	logger         plog.Logger
	queryerContext sqlx.QueryerContext
}

func (qcl *queryContextLogger) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	qcl.logger.Query(ctx, query, args)
	return qcl.queryerContext.QueryContext(ctx, query, args...)
}

func (qcl *queryContextLogger) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	qcl.logger.Query(ctx, query, args)
	return qcl.queryerContext.QueryxContext(ctx, query, args...)
}

func (qcl *queryContextLogger) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	qcl.logger.Query(ctx, query, args)
	return qcl.queryerContext.QueryRowxContext(ctx, query, args...)
}
