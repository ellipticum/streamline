package services

import (
	"context"
	"github.com/ellipticum/streamline/internal/dto"
	"github.com/ellipticum/streamline/internal/models"
	"github.com/ellipticum/streamline/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type User struct{}

var delay = 10 * time.Second

func (user User) Get() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), delay)

	defer cancel()

	cursor, err := storage.Client.Database("streamline").Collection("users").Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var results []models.User

	for cursor.Next(ctx) {
		var elem models.User

		err := cursor.Decode(&elem)

		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (user User) Create(data dto.UserCreateDTO) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), delay)

	defer cancel()

	result, err := storage.Client.Database("streamline").Collection("users").InsertOne(ctx, data)

	if err != nil {
		return nil, err
	}

	return result, nil
}
