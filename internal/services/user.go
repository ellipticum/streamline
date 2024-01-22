package services

import (
	"context"
	"github.com/ellipticum/streamline/internal/dto"
	"github.com/ellipticum/streamline/pkg/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type User struct{}

var delay = 10 * time.Second

func (user User) Get() (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), delay)

	defer cancel()

	result, err := storage.Client.Database("streamline").Collection("users").Find(ctx, nil)

	if err != nil {
		return nil, err
	}

	return result, nil
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
