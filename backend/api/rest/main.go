package main

import (
	"flag"
	"log"
	"os"

	"github.com/briand787b/piqlit/api/rest/controller"
	"github.com/briand787b/piqlit/core/fs"
	"github.com/briand787b/piqlit/core/plog"
	"github.com/briand787b/piqlit/core/postgres"

	"github.com/google/uuid"
)

var (
	portFlag    = flag.Int("port", 0, "the port to listen on")
	dataDirFlag = flag.String("data", "", "directory to look for data")
)

func main() {
	flag.Parse()

	l := plog.NewPLogger(log.New(os.Stdout, "", 0), uuid.New())
	ms := postgres.NewMediaPGStore(l, postgres.GetExtFull(l))
	os := fs.NewObjectFileStore(l, *dataDirFlag)

	controller.Serve(*portFlag, l, ms, os)
}
