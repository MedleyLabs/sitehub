// Package sitehub provides a federated hub for model versioning.
// It allows users to download / upload machine learning models.
package main

import (
	"encoding/json"
	"fmt"
	//"log"
	"net/http"
	//"time"

	"github.com/gorilla/mux"
)

//func main() {
//	startFileServer()
//}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	fmt.Println(decoder)
	fmt.Println(r)
	vars := mux.Vars(r)
	fmt.Println(vars)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Patient ID: %v\n", vars["patient_id"])
}

func ModelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

//func main() {
//	r := mux.NewRouter()
//	r.HandleFunc("/", HomeHandler)
//	r.HandleFunc("/data", DataHandler)
//	r.HandleFunc("/model", ModelHandler)
//	http.Handle("/", r)
//
//	srv := &http.Server{
//		Handler: r,
//		Addr:    "127.0.0.1:8000",
//		WriteTimeout: 15 * time.Second,
//		ReadTimeout:  15 * time.Second,
//	}
//
//	log.Fatal(srv.ListenAndServe())
//}

func main() {
	startAPI()
}
