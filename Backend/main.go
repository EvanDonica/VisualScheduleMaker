package main

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Visual Schedule Maker"))
}

func main() {
	http.HandleFunc("/home", homeHandler)
	http.ListenAndServe(":8080", nil)
}
