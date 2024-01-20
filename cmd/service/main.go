package main

import (
	"fmt"
	"github.com/ellipticum/streamline/core/router"
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

	address := ":" + port

	routes.Register()

	http.HandleFunc("/", router.Listen)
	http.ListenAndServe(address, nil)
}
