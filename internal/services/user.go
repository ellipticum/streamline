package services

import (
	"context"
	"fmt"
	"github.com/ellipticum/streamline/internal/dto"
	"github.com/ellipticum/streamline/internal/models"
	"github.com/ellipticum/streamline/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type User struct{}

var delay = 10 * time.Second

func (user User) Get() ([]*models.User, error) {
	result, err := storage.Find("streamline", "users", &models.User{}, bson.D{}, delay)

	if err != nil {
		return nil, err
	}

	data, ok := result.([]*models.User)

	if !ok {
		return nil, fmt.Errorf("Could not cast result to []*models.User")
	}

	return data, nil
}

func (user User) GetByID(ID string) (*models.User, error) {
	target, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": target}

	result, err := storage.Find("streamline", "users", &models.User{}, filter, delay)

	if err != nil {
		return nil, err
	}

	data, ok := result.([]*models.User)

	if len(data) == 0 {
		return nil, fmt.Errorf("Could not find user with ID %s", ID)
	}

	if !ok {
		return nil, fmt.Errorf("Could not cast result to *models.User")
	}

	return data[0], nil
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
