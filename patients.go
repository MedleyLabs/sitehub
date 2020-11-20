package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Patient struct {
	ID        string `json:"id"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

var patients = []Patient{}
var idCounter int

func startAPI() {
	r := mux.NewRouter()
	patientsR := r.PathPrefix("/patients").Subrouter()
	patientsR.Path("").Methods(http.MethodGet).HandlerFunc(getAllPatients)
	patientsR.Path("").Methods(http.MethodPost).HandlerFunc(createPatient)
	patientsR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(getPatientByID)
	patientsR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(updatePatient)
	patientsR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(deletePatient)

	fmt.Println("Start listening")
	fmt.Println(http.ListenAndServe(":8080", r))
}

func getAllPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(patients); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getPatientByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not implemented")
}

func updatePatient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not implemented")
}

func deletePatient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented")
}

func createPatient(w http.ResponseWriter, r *http.Request) {
	p := Patient{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	patients = append(patients, p)
	response, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
