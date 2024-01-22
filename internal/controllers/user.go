package controllers

import (
	"encoding/json"
	"github.com/ellipticum/streamline/internal/dto"
	"github.com/ellipticum/streamline/internal/services"
	"github.com/ellipticum/streamline/pkg/utils"
	"net/http"
)

var service = services.User{}

type User struct{}

func (user *User) Get(w http.ResponseWriter, r *http.Request) {
	result, err := service.Get()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)
}

func (user *User) GetByID(w http.ResponseWriter, r *http.Request) {
	ID := utils.Extract(r)

	if ID == "" {
		http.Error(w, "No ID provided", http.StatusBadRequest)
		return
	}

	result, err := service.GetByID(ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)
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

	result, err := service.Create(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)
}
