package main

import (
	"fmt"
	"github.com/ellipticum/streamline/core/router/mux"
	"github.com/ellipticum/streamline/internal/routes"
	"github.com/ellipticum/streamline/pkg/utils"
	"net/http"
	"os"
)

func main() {
	err := utils.LoadEnvFromFile("../../env.development")

	if err != nil {
		fmt.Printf("Error loading .env: %s\n", err)

		return
	}

	port := os.Getenv("PORT")

	mux := mux.Create()

	routes.Register(mux)

	address := ":" + port

	fmt.Println(address)

	http.ListenAndServe(address, mux)

	fmt.Printf("The server has been started. Port: %s\n", port)
}
