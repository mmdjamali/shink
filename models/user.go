package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct to describe User object.
type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
