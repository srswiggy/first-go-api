package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/srswiggy/first-api/requests"
	"github.com/srswiggy/first-api/role"
	"github.com/srswiggy/first-api/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Request for CreateUser")

	var userRequest requests.UserRequest

	userRole := role.GetRole(r.Header.Values("role")[0])

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		log.Print("error decoding body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newUser *user.User
	if userRole == role.Admin {
		newUser = user.CreateAmdin(userRequest.Name, userRequest.Email)
	} else if userRole == role.SpecialUser {
		newUser = user.CreateSpecialUser(userRequest.Name, userRequest.Email)
	} else {
		newUser = user.CreateUser(userRequest.Name, userRequest.Email)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&newUser); err != nil {
		log.Printf("error encoding body response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Print("user created successfully")
}
