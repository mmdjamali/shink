package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Otp struct {
	ID                    primitive.ObjectID `json:"_id" bson:"_id"`
	Email                 string             `json:"email" bson:"email"`
	Digits                string             `json:"digits" bson:"digits"`
	CreatedAt             time.Time          `json:"created_at" bson:"created_at"`
	AttemptCount          int                `json:"attempt_count" bson:"attempt_count"`
	UpdatedAttemptCountAt time.Time          `json:"updated_attempt_count_at" bson:"updated_attempt_count_at"`
	UpdatedDigitsAt       time.Time          `json:"updated_digits_at" bson:"updated_digits_at"`
}
