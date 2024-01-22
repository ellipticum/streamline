package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	Client *mongo.Client
	err    error
)

func Connect(address string, attempts int) {
	delay := 10 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), delay)

	defer cancel()

	for i := 0; i < attempts; i++ {
		Client, err = mongo.Connect(ctx, options.Client().ApplyURI(address))
	}

	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}
}
