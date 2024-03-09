package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/form", FormHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	// TODO read data from the request
	// TODO save the data
	// TODO respond to the request Excercise for you
}
