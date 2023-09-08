package services

import (
	"context"
	"fmt"
	"start/config"
	"start/database"
	"start/models"
	"start/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OtpService struct {
	Email string `json:"email" bson:"email"`
}

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
		ID:                    primitive.NewObjectID(),
		Email:                 otpS.Email,
		Digits:                "444444",
		CreatedAt:             time.Now(),
		UpdatedAttemptCountAt: time.Now(),
		AttemptCount:          1,
		UpdatedDigitsAt:       time.Now(),
	})

	if err != nil {
		return "", err
	}

	return "444444", nil
}

func (otpS *OtpService) Update() (string, error) {
	otpCollection := database.DB.Collection("otps")

	otp, otp_err := otpS.Get()

	if otp_err != nil {
		return "", otp_err
	}

	can, can_err := otpS.CanCreateNew(otp)

	if can_err != nil {
		return "", can_err
	}

	if !can {
		if err := otpS.ResetAttempts(otp); err != nil {
			return "", err
		}
	}

	_, err := otpCollection.UpdateOne(
		context.TODO(),
		fiber.Map{
			"email": otpS.Email,
		}, bson.M{
			"$set": bson.M{
				"updated_digits_at": time.Now(),
				"digits":            "444444",
			},
			"$inc": bson.M{
				"attempt_count": 1,
			},
		})

	if err != nil {
		fmt.Println(err.Error())

		return "", err
	}

	return "444444", nil
}

func (otpS *OtpService) Validate(d string) (bool, error) {
	otp, err := otpS.Get()

	if time.Now().After(otp.UpdatedDigitsAt.Add(config.OtpLifeTime)) {
		return false, &utils.CustomError{Message: "otp has been expired"}
	}

	if err != nil {
		return false, err
	}

	return d == otp.Digits, nil
}

func (otpS *OtpService) Get() (*models.Otp, error) {
	otpCollection := database.DB.Collection("otps")

	otp := models.Otp{}

	res := otpCollection.FindOne(context.TODO(), fiber.Map{
		"email": otpS.Email,
	})

	if res.Err() != nil {
		return nil, res.Err()
	}

	if err := res.Decode(&otp); err != nil {

	}

	return &otp, nil
}

func (otpS *OtpService) CanCreateNew(otp *models.Otp) (bool, error) {
	if otp.AttemptCount < 15 {
		return true, nil
	}

	fmt.Println("can't")

	return false, nil
}

func (otpS *OtpService) ResetAttempts(otp *models.Otp) error {
	otpCollection := database.DB.Collection("otps")

	if !time.Now().After(otp.UpdatedAttemptCountAt.Add(time.Hour * 24)) {
		return &utils.CustomError{Message: "request limit for otp"}
	}

	_, err := otpCollection.UpdateOne(context.TODO(), bson.M{
		"_id": otp.ID,
	}, bson.M{
		"$set": bson.M{
			"attempt_count":            0,
			"updated_attempt_count_at": time.Now(),
		},
	})

	if err != nil {
		return err
	}

	return nil
}
