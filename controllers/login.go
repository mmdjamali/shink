package controllers

import (
	"fmt"
	"start/services"
	"start/utils"

	"github.com/gofiber/fiber/v2"
)

type Credentials struct {
	Email string `json:"email" bson:"email"`
}

func LoginController(c *fiber.Ctx) error {
	body := &Credentials{}

	if err := c.BodyParser(body); err != nil {
		return err
	}

	if body.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
		})
	}

	if !utils.IsValidEmail(body.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"message": "your email address is not valid!",
		})
	}

	userS := &services.UserService{Email: body.Email}
	otpS := &services.OtpService{Email: body.Email}
	// otps := database.DB.Collection("otps")

	exists, err := userS.Exists()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"message": "something went wrong while checking user does exists",
		})
	}

	if !exists {
		_, user_err := userS.Create()

		if user_err != nil {
			return user_err
		}

		_, otp_err := otpS.Create()

		if otp_err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"data":    nil,
				"message": "we were not able to create a otp, please try again!",
			})
		}

		return c.Redirect(fmt.Sprintf("/confirm?email=%s", body.Email))
	}

	_, get_user_err := userS.Get()

	if get_user_err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"message": "something went wrong while getting user",
		})
	}

	// rand.Seed(time.Now().UnixMilli())

	// new_otp := rand.Intn(999999)
	// new_otp := fmt.Sprint(444444)

	_, update_err := otpS.Update()

	if update_err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"message": "we were not able to create a one time password for you, please try again!",
		})
	}

	return c.Redirect(fmt.Sprintf("/confirm?email=%s", body.Email))

}
