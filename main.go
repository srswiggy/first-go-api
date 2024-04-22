package main

import (
	"encoding/json"
	"fmt"
	mux2 "github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	mux := mux2.NewRouter()
	mux.HandleFunc("/user", postBody).Methods("POST")

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

type User struct {
	Name  string
	Email string
}

func postBody(w http.ResponseWriter, r *http.Request) {
	var usr User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		fmt.Print("error decoding body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usr)
}
