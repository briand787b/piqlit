package postgrestest

import (
	"log"
	"testing"
	"time"

	"github.com/briand787b/piqlit/core/plog"
	"github.com/briand787b/piqlit/core/postgres"
	"github.com/briand787b/piqlit/core/psql"
	"github.com/briand787b/piqlit/core/test"
)

// PGHelper is a Helper that specifically helps
// with cleaning up after PG-interacting tests
type PGHelper struct {
	test.Helper
	DB        psql.ExtFull
	ParentIDs map[Table]int
}

// NewPGHelper returns a new PGHelper with all
// necessary setup/connection operations complete
func NewPGHelper(t *testing.T) *PGHelper {
	tc := test.SetTimeout(5 * time.Second)
	defer func() { tc <- struct{}{} }()

	l := &plog.SOLogger{}
	return &PGHelper{
		Helper: test.Helper{
			T:  t,
			L:  l,
			Tm: time.Now().UTC().Truncate(time.Second),
			CF: test.NewCleaner(func() { log.Println("Postgres Cleaned!!!!") }),
		},
		DB:        postgres.GetExtFull(l),
		ParentIDs: make(map[Table]int),
	}
}
