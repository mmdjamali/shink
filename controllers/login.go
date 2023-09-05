package controllers

import (
	"context"
	"fmt"
	"math/rand"
	"start/models"
	"start/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Credentials struct {
	Email string `json:"email" bson:"email"`
}

func LoginController(c *fiber.Ctx) error {
	credential := &Credentials{}

	if err := c.BodyParser(credential); err != nil {
		return err
	}

	if credential.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": "false",
			"data":    nil,
		})
	}

	db := utils.ConnectDB().Database("test-go")

	users := db.Collection("users")
	otps := db.Collection("otps")

	filter := bson.M{"email": credential.Email}

	result := users.FindOne(context.TODO(), filter)

	if result.Err() != nil {
		user, user_err := users.InsertOne(context.TODO(), bson.M{
			"email":      credential.Email,
			"created_at": time.Now(),
		})

		if user_err != nil {
			return user_err
		}

		new_otp := fmt.Sprint(444444)

		fmt.Println(user.InsertedID.(primitive.ObjectID))

		_, otp_err := otps.InsertOne(context.TODO(), models.Otp{
			Email:  credential.Email,
			Digits: new_otp,
		})

		// fmt.Printf(otp.InsertedID.(string))

		if otp_err != nil {
			return otp_err
		}

		return c.Redirect(fmt.Sprintf("/confirm?email=%s", credential.Email))
	}

	user := &models.User{}

	if err := result.Decode(user); err != nil {
		return err
	}

	rand.Seed(time.Now().UnixMilli())

	// new_otp := rand.Intn(999999)
	new_otp := fmt.Sprint(444444)

	_, err := otps.InsertOne(context.TODO(), models.Otp{
		Email:     credential.Email,
		Digits:    new_otp,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	return c.Redirect(fmt.Sprintf("/confirm?email=%s", credential.Email))

}
