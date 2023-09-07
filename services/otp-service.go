package services

import (
	"context"
	"fmt"
	"start/database"
	"start/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OtpService struct {
	Email string `json:"email" bson:"email"`
}

// func (otpS OtpService) New(e string) *OtpService {
// 	return &OtpService{
// 		Email: e,
// 	}
// }

func (otpS *OtpService) Exists() (bool, error) {
	otpCollection := database.DB.Collection("otps")

	res := otpCollection.FindOne(context.TODO(), fiber.Map{
		"email": otpS.Email,
	})

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, res.Err()
	}
	return true, nil
}

func (otpS *OtpService) Create() (string, error) {
	otpCollection := database.DB.Collection("otps")

	_, err := otpCollection.InsertOne(context.TODO(), models.Otp{
		Email:     otpS.Email,
		Digits:    "444444",
		CreatedAt: time.Now(),
	})

	if err != nil {
		return "", err
	}

	return "444444", nil
}

func (otpS *OtpService) Update() (string, error) {
	otpCollection := database.DB.Collection("otps")

	_, err := otpCollection.UpdateOne(
		context.TODO(),
		fiber.Map{
			"email": otpS.Email,
		}, bson.M{
			"$set": bson.M{
				"digits": "444444",
			},
		})

	if err != nil {
		fmt.Println(err.Error())

		return "", err
	}

	return "444444", nil
}

func (otpS *OtpService) Validate(d string) (bool, error) {
	otpCollection := database.DB.Collection("otps")

	otp := models.Otp{}

	res := otpCollection.FindOne(context.TODO(), fiber.Map{
		"email": otpS.Email,
	})

	if res.Err() != nil {
		return false, res.Err()
	}

	if err := res.Decode(&otp); err != nil {

	}

	return d == otp.Digits, nil
}

func (otpS *OtpService) Get(d string) (bool, error) {
	otpCollection := database.DB.Collection("otps")

	otp := models.Otp{}

	res := otpCollection.FindOne(context.TODO(), fiber.Map{
		"email": otpS.Email,
	})

	if res.Err() != nil {
		return false, res.Err()
	}

	if err := res.Decode(&otp); err != nil {

	}

	return d == otp.Digits, nil
}
