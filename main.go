package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func handler() {
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

	app.Listen(":3001")
}
