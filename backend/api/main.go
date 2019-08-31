package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	phttp "github.com/briand787b/piqlit/api/http"
)

var (
	portFlag   = flag.Int("port", 8080, "the port to listen on")
	masterFlag = flag.Bool("master", false, "server becomes master if true")
)

func main() {
	flag.Parse()
	env := os.Getenv("ENV")
	if env == "" {
		env = "defualt_env"
	}

	httpServer := phttp.NewChiServer(*masterFlag, &phttp.ServerArgs{})

	status := "SLAVE"
	if *masterFlag {
		status = "MASTER"
	}

	log.Printf("Starting %s HTTP server on port %v...\n", status, *portFlag)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%v", *portFlag), httpServer))
}
