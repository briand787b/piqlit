package postgres_test

import (
	"flag"
	"os"
	"testing"
)

var live = flag.Bool("live", false, "use live dependencies - pkg dependent")

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func skipNotLive(t *testing.T) {
	if !*live {
		t.Skip()
	}
}
