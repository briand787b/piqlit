package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/briand787b/piqlit/core/model"
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

	fmt.Printf("%v\n", model.Server{})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, from " + env))
	})

	status := "SLAVE"
	if *masterFlag {
		status = "MASTER"
	}

	log.Printf("STARTING %s SERVER ON PORT %v...\n", status, *portFlag)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%v", *portFlag), nil))
}
