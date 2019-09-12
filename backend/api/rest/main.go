package main

import (
	"flag"

	"github.com/briand787b/piqlit/api/rest/controller"
	"github.com/briand787b/piqlit/core/obj"

	"github.com/sirupsen/logrus"
)

var (
	portFlag    = flag.Int("port", 8080, "the port to listen on")
	masterFlag  = flag.Bool("master", false, "server becomes master if true")
	dataDirFlag = flag.String("data", "", "directory to look for data")
)

func main() {
	flag.Parse()

	l := logrus.New()
	// l.Formatter = &logrus.JSONFormatter{}

	os := obj.NewObjectFileStore(l, *dataDirFlag)

	controller.Serve(*portFlag, l, os)
}
