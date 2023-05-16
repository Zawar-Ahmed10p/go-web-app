package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "hello world")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("No of bytes: ", n)
	})
	http.ListenAndServe(":8080", nil)
}
