package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Owner          primitive.ObjectID `json:"owner" bson:"owner"`
	Link           string             `json:"link" bson:"link"`
	RedirectCount  int                `json:"redirect_count" bson:"redirect_count"`
	Custom         string             `json:"custom" bson:"custom"`
	LastRedirectAt time.Time          `json:"last_redirect_at" bson:"last_redirect_at"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
}
