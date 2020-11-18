package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// The directory from which to serve files.
const Dir = "."

// The address from which to serve files.
const Addr = ":9000"

// startFileServer starts the file server to download files.
func startFileServer() {
	fs := http.FileServer(http.Dir(Dir))

	index, err := ioutil.ReadFile("index.html")

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprint(res, string(index))
	})

	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Fatal(http.ListenAndServe(Addr, nil))
}
