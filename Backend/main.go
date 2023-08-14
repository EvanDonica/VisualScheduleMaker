package main

import (
	"backend/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)

	http.Handle("/", r)

	// Replace 8080 with the desired port number
	http.ListenAndServe(":8080", nil)
}
