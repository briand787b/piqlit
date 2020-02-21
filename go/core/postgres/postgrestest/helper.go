package postgrestest

import (
	"context"
	"flag"
	"log"
	"os"
	"testing"
	"time"

	"github.com/briand787b/piqlit/core/plog/plogtest"
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
	flag.Parse()
	if testing.Short() {
		log.Println("skipping tests because `-test.short` flag provided")
		os.Exit(0)
	}

	tc := test.SetTimeout(5 * time.Second)
	defer func() { tc <- struct{}{} }()

	l := plogtest.NewMockLogger()
	db := postgres.GetExtFull(l)

	tx, err := db.Begin(context.Background(), nil)
	if err != nil {
		t.Fatal("NewPGHelper could not begin tx: ", err)
	}

	return &PGHelper{
		Helper: test.Helper{
			T:  t,
			L:  l,
			Tm: time.Now().UTC().Truncate(time.Second),
			CF: test.NewCleaner(func() {
				if err := tx.Rollback(); err != nil {
					t.Fatal("could not roll back tx: ", err)
				}

				log.Println("Postgres Cleaned!!!!")
			}),
		},
		DB:        psql.NewExtFullFromTx(l, tx),
		ParentIDs: make(map[Table]int),
	}
}

// ParentID returns the ParentID for the given Table
func (h *PGHelper) ParentID(t Table) *int {
	id, ok := h.ParentIDs[t]
	if !ok {
		defer h.Clean()
		h.T.Fatalf("could not find parentID from table %v", t)
	}

	return &id
}
