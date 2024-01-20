package routes

import (
	"encoding/json"
	"fmt"
	"github.com/ellipticum/streamline/core/router"
	"net/http"
)

type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Register() {
	router.Get("/api/v1/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	router.Post("/api/v1/test", func(w http.ResponseWriter, r *http.Request) {
		var data Data

		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		fmt.Fprintln(w, data)
	})
}
