package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID        string `json:"id"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

var users = []User{}
var idCounter int

func startAPI() {
	r := mux.NewRouter()
	usersR := r.PathPrefix("/users").Subrouter()
	usersR.Path("").Methods(http.MethodGet).HandlerFunc(getAllUsers)
	usersR.Path("").Methods(http.MethodPost).HandlerFunc(createUser)
	usersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(getUserByID)
	usersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(updateUser)
	usersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(deleteUser)

	fmt.Println("Start listening")
	fmt.Println(http.ListenAndServe(":8080", r))
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not implemented")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not implemented")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	p := User{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users = append(users, p)
	response, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
