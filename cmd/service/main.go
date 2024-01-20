package main

import (
	"fmt"
	"github.com/ellipticum/streamline/core/router/mux"
	"github.com/ellipticum/streamline/internal/routes"
	"net/http"
)

func main() {
	mux := mux.Create()

	routes.Register(mux)

	http.ListenAndServe(":8080", mux)

	fmt.Println("The server has been started")
}
