package services

import (
	"context"
	"start/database"
	"start/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LinkService struct {
	Custom string `json:"custom" bson:"custom"`
	Link   string `json:"digits" bson:"digits"`
}

func (ls *LinkService) Exists() (bool, error) {
	linkCollection := database.DB.Collection("links")

	res := linkCollection.FindOne(context.TODO(), fiber.Map{
		"custom": ls.Custom,
	})

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, res.Err()
	}
	return true, nil
}

func (ls *LinkService) Create(owner string) error {
	linkCollection := database.DB.Collection("links")

	owner_id, owner_id_err := primitive.ObjectIDFromHex(owner)

	if owner_id_err != nil {
		return owner_id_err
	}

	_, err := linkCollection.InsertOne(context.TODO(), models.Link{
		ID:             primitive.NewObjectID(),
		Owner:          owner_id,
		Link:           ls.Link,
		Custom:         ls.Custom,
		RedirectCount:  0,
		LastRedirectAt: time.Now(),
		CreatedAt:      time.Now(),
	})

	if err != nil {
		return err
	}
	return nil
}

func (ls *LinkService) Get() (*models.Link, error) {
	linkCollection := database.DB.Collection("links")

	link := &models.Link{}

	res := linkCollection.FindOne(context.TODO(), bson.M{
		"custom": ls.Custom,
	})

	if res.Err() != nil {
		return nil, res.Err()
	}

	if err := res.Decode(&link); err != nil {
		return nil, err
	}

	return link, nil
}

func (ls *LinkService) UpdateRedirectCount() error {
	linkCollection := database.DB.Collection("links")

	_, err := linkCollection.UpdateOne(context.TODO(), bson.M{
		"custom": ls.Custom,
	}, bson.M{
		"$inc": bson.M{
			"redirect_count": 1,
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func (ls *LinkService) GetLinksOfUser(_id string) (*[]models.Link, error) {
	linkCollection := database.DB.Collection("links")

	owner_id, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		return nil, err
	}

	var links []models.Link

	res, err := linkCollection.Find(context.TODO(), bson.M{
		"owner": owner_id,
	})

	if err != nil {
		return nil, err
	}

	if err := res.All(context.TODO(), &links); err != nil {
		return nil, err
	}

	return &links, err

}
