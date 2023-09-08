package controllers

import (
	"start/services"

	"github.com/gofiber/fiber/v2"
)

func RedirectController(c *fiber.Ctx) error {
	custom := c.Params("custom")

	if custom == "" {
		return c.Redirect("/home")
	}

	LS := services.LinkService{
		Custom: custom,
	}

	link, err := LS.Get()

	if err != nil {
		return c.Redirect("/home")
	}

	update_err := LS.UpdateRedirectCount()

	if update_err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
		})
	}

	return c.Redirect(link.Link)
}
