package obj_test

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/briand787b/piqlit/core/obj"

	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
)

var update = flag.Bool("update", false, "update test fixture with generated data")

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestObjectFileStoreList(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		expErrIsNil bool
		expList     []string
	}{
		{"nonexistant_dir_does_not_error", true, nil},
		{"empty_dir_returns_nil", true, nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actList, err := obj.NewObjectFileStore(
				logrus.New(),
				"./testdata/"+t.Name(),
			).ListNames(context.Background())

			if !tt.expErrIsNil {
				if err == nil {
					t.Fatal("expected err to be non-nil, was nil")
				}

				return
			}

			if err != nil {
				t.Fatal("expected err to be nil, was ", err)
			}

			if !cmp.Equal(actList, tt.expList) {
				t.Fatalf("expected list and actual list are different: %s",
					cmp.Diff(actList, tt.expList),
				)
			}
		})
	}
}
