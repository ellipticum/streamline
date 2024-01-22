package main

import (
	"context"
	"fmt"
	"github.com/ellipticum/streamline/core/router"
	"github.com/ellipticum/streamline/internal/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")

	delay := 10 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), delay)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongo:27017"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	address := ":" + port

	routes.Register()

	fmt.Printf("The server has been started. Port | %s\n", port)

	http.HandleFunc("/", router.Listen)
	http.ListenAndServe(address, nil)
}
