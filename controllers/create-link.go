package controllers

import (
	"start/services"
	"start/sessions"
	"start/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateLink(c *fiber.Ctx) error {
	body := &struct {
		Link   string `json:"link" bson:"link"`
		Custom string `json:"custom" bson:"custom"`
	}{}

	if err := c.BodyParser(body); err != nil {
		return err
	}

	if !utils.IsValidURL(body.Link) {
		return c.Redirect("/")
	}

	if !utils.IsDomainDiffrent(body.Link) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "aren't you so smart? :)",
		})
	}

	sess, sess_err := sessions.SessionStore.Get(c)

	if sess_err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "faild to get token",
		})
	}

	uid := sess.Get("uid")

	if uid == nil {
		return c.Redirect("/")
	}

	defer sess.Save()

	LS := &services.LinkService{
		Custom: body.Custom,
		Link:   body.Link,
	}

	exists, exists_err := LS.Exists()

	if exists_err != nil {
		return c.Redirect("/")
	}

	if exists {
		return c.Redirect("/")
	}

	cr_err := LS.Create(uid.(string))

	if cr_err != nil {
		return c.Redirect("/")
	}

	return c.Redirect("/app")
}
