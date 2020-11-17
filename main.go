package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	dir := "."
	addr := ":9000"

	fs := http.FileServer(http.Dir(dir))

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprint(res, "<h1>Sitehub</h1>")
	})

	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Fatal(http.ListenAndServe(addr, nil))
}
