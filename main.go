package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Use(logger.New())
	app.Use(cors.New())

	app.Static("/public", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("main", fiber.Map{
			"Title": "Shink - Free Link shortener",
		})
	})

	app.Get("/auth", func(c *fiber.Ctx) error {
		return c.Render("pages/auth", fiber.Map{
			"Title": "Shink - Authentication",
		})
	})

	app.Get("/confirm", func(c *fiber.Ctx) error {
		fmt.Println(c.Query("email"))
		return c.Render("pages/confirm-email", fiber.Map{
			"Title": "Shink - Authentication",
			"Email": c.Query("email"),
		})
	})

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/login", func(c *fiber.Ctx) error {
		var payload struct {
			Email string `json:"email"`
		}

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"success": false,
				"data":    nil,
			})
		}

		return c.Redirect(fmt.Sprintf("/confirm?email=%s", payload.Email))
	})

	v1.Post("/confirm", func(c *fiber.Ctx) error {
		var payload struct {
			Digits string `json:"digits"`
			Email  string `json:"email"`
		}

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"success": false,
				"data":    nil,
			})
		}

		fmt.Println(payload)

		return c.Redirect("/")
	})

	app.Listen(":3001")
}
