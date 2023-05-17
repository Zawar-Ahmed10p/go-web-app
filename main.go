package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about page")
}
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	log.Printf("Starting server on %+s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
