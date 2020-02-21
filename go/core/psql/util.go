package psql

import (
	"context"

	"github.com/briand787b/piqlit/core/plog"
)

func logQuery(l plog.Logger, qry string, args ...interface{}) {
	l.Query(context.Background(), qry, args...)
}
