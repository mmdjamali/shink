package main

import (
	"fmt"
	"start/config"
	"start/controllers"
	"start/database"
	"start/sessions"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func init() {
	// load env before everything
	config.LoadEnv()

	// connect to database
	database.DatabaseInit()

	//initilaize session storage
	sessions.InitilaizeSessionStore()
}

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
		sess, sess_err := sessions.SessionStore.Get(c)

		if sess_err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "faild to get token",
			})
		}

		uid := sess.Get("uid")
		defer sess.Save()

		fmt.Println(uid)

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
