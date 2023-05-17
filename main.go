package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}
func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		fmt.Println("Parsing error", err)
		return
	}
	err = parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Parsing error", err)
		return
	}
}
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	log.Printf("Starting server on %+s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
