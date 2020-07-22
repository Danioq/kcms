package main

import (
	"github.com/gorilla/mux"

	"fmt"
	"net/http"
)

//Index - Hello world
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World")
}

//HelloNumber - get number
func HelloNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, id)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/{id}", HelloNumber)

	go http.ListenAndServe(":8080", router)
	go http.ListenAndServe(":443", router)
}
