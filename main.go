package main

import (
	"fmt"
	"start/controllers"

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

	v1.Post("/login", controllers.LoginController)

	v1.Post("/confirm", controllers.ValidateOtp)

	app.Listen(":3001")
}
