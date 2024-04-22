package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my API\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
