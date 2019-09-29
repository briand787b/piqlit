package psql

import "github.com/briand787b/piqlit/core/plog"

func logQuery(l plog.Logger, qry string, args ...interface{}) {
	l.Infow("[PG QUERY]",
		"query", qry,
		"args", args,
	)
}
