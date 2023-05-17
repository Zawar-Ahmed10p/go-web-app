package main

import (
	"log"
	"net/http"

	"github.com/Zawar-Ahmed10p/go-web-app/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	log.Printf("Starting server on %+s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
