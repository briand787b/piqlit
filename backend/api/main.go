package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "defualt_env"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, from " + env))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
