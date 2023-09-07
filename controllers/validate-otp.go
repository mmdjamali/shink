package controllers

import (
	"start/services"
	"start/sessions"

	"github.com/gofiber/fiber/v2"
)

func ValidateOtp(c *fiber.Ctx) error {
	body := struct {
		Email  string `json:"email" bson:"email"`
		Digits string `json:"digits" bson:"digits"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	userS := &services.UserService{Email: body.Email}
	otpS := &services.OtpService{Email: body.Email}

	exists, exists_err := userS.Exists()

	if exists_err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "something went wrong!",
		})
	}

	if !exists {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "something went wrong!",
		})
	}

	valid, valid_err := otpS.Validate(body.Digits)

	if valid_err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
		})
	}

	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
		})
	}

	user, user_err := userS.Get()

	if user_err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": user_err.Error(),
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
