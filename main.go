package main

import (
	"fmt"
	"log"
	"net/http" //responsible for basic http server
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Visagan")
	})
	http.HandleFunc("/reg", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376211SE151")
	})
	http.HandleFunc("/dept", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Information Science and Engineering")
	})
	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Blue")
	})
	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
