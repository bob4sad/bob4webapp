package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHendler).Methods("GET")
	r.HandleFunc("/goodbye", goodbyeHendler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func helloHendler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func goodbyeHendler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Goodbye, world!")
}