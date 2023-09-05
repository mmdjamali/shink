package controllers

import (
	"context"
	"start/database"
	"start/models"
	"start/sessions"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func ValidateOtp(c *fiber.Ctx) error {
	credentials := struct {
		Email  string `json:"email" bson:"email"`
		Digits string `json:"digits" bson:"digits"`
	}{}

	// ctx, ctx_err := context.WithTimeout(context.TODO(), 10*time.Second)

	// if ctx_err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "context error",
	// 	})
	// }

	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	otps := database.DB.Collection("otps")

	otp := new(models.Otp)

	if err := otps.FindOne(context.TODO(), fiber.Map{}).Decode(&otp); err != nil {
		return err
	}

	if otp.Digits != credentials.Digits {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
		})
	}

	var user models.User
	res := database.DB.Collection("users").FindOne(context.TODO(), bson.M{
		"email": credentials.Email,
	})

	if res.Err() != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": res.Err().Error(),
		})
	}

	if d_err := res.Decode(&user); d_err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": d_err.Error(),
		})
	}

	sess, sess_err := sessions.SessionStore.Get(c)

	if sess_err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": sess_err.Error(),
		})
	}

	sess.Set("uid", user.ID.Hex())

	if save_err := sess.Save(); save_err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": save_err.Error(),
		})
	}

	return c.Redirect("/")
}
