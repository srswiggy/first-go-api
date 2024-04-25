package main

import (
	"log"
	"net/http"

	mux2 "github.com/gorilla/mux"
	"github.com/srswiggy/first-api/controllers"
)

func main() {
	mux := mux2.NewRouter()
	mux.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

