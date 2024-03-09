package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
