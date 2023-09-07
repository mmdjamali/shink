package models

import (
	"time"
)

type Otp struct {
	Email     string    `json:"email" bson:"email"`
	Digits    string    `json:"digits" bson:"digits"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
