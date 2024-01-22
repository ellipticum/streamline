package main

import (
	"fmt"
	"github.com/ellipticum/streamline/core/router"
	"github.com/ellipticum/streamline/internal/routes"
	"github.com/ellipticum/streamline/pkg/storage"
	"net/http"
	"os"
)

func main() {
	storage.Connect("mongodb://root:root@mongo:27017", 10)

	port := os.Getenv("PORT")

	address := ":" + port

	routes.Register()

	fmt.Printf("The server has been started. Port | %s\n", port)

	http.HandleFunc("/", router.Listen)
	http.ListenAndServe(address, nil)
}
