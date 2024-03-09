package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type FormData struct {
	Year       string `json:"year"`
	CourseName string `json:"course_name"`
	Link       string `json:"file_link"`
}

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
	var data = FormData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Invalid Form data", 400)
		return
	}

	// TODO save the data
	// TODO respond to the request Excercise for you
}
