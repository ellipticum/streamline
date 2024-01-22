package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ellipticum/streamline/internal/dto"
	"github.com/ellipticum/streamline/internal/services"
	"net/http"
)

var service = services.User{}

type User struct{}

func (user *User) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User | Get")
}

func (user *User) Create(w http.ResponseWriter, r *http.Request) {
	var data dto.UserCreateDTO

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.Email == "" || data.Username == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	result := service.Create(data)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)
}
