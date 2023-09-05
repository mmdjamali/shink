package controllers

import (
	"context"
	"start/models"
	"start/utils"

	"github.com/gofiber/fiber/v2"
)

func ValidateOtp(c *fiber.Ctx) error {
	credentials := struct {
		Email  string `json:"email" bson:"email"`
		Digits string `json:"digits" bson:"digits"`
	}{}

	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
		})
	}

	otps := utils.ConnectDB().Database("test-go").Collection("otps")

	otp := new(models.Otp)

	if err := otps.FindOne(context.TODO(), fiber.Map{}).Decode(&otp); err != nil {
		return err
	}

	if otp.Digits == credentials.Digits {
		return c.Redirect("/")
	}

	return c.JSON(fiber.Map{})
}
