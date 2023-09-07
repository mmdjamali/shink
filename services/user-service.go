package services

import (
	"context"
	"start/database"
	"start/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	Email string `json:"email" bson:"email"`
}

func (userS *UserService) Exists() (bool, error) {
	usersCollection := database.DB.Collection("users")

	res := usersCollection.FindOne(context.TODO(), bson.M{
		"email": userS.Email,
	})

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, res.Err()
	}
	return true, nil
}

func (userS *UserService) Create() (string, error) {
	usersCollection := database.DB.Collection("users")

	res, err := usersCollection.InsertOne(context.TODO(), models.User{
		Email:     userS.Email,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return "", err
	}

	return res.InsertedID.(string), nil
}

func (userS *UserService) Get() (*models.User, error) {
	user := models.User{}
	usersCollection := database.DB.Collection("users")

	res := usersCollection.FindOne(context.TODO(), bson.M{
		"email": userS.Email,
	})

	if res.Err() != nil {
		return nil, res.Err()
	}

	if err := res.Decode(user); err != nil {
		return nil, err
	}

	return &user, nil
}
