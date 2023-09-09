package main

import (
	"fmt"
	"start/config"
	"start/controllers"
	"start/database"
	"start/services"
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

	app.Get("/home", func(c *fiber.Ctx) error {
		sess, sess_err := sessions.SessionStore.Get(c)

		if sess_err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "faild to get token",
			})
		}

		uid := sess.Get("uid")
		defer sess.Save()

		return c.Render("main", fiber.Map{
			"Title": "Shink - Free Link shortener",
			"UID":   uid,
		})
	})

	app.Get("/app", func(c *fiber.Ctx) error {
		sess, sess_err := sessions.SessionStore.Get(c)

		if sess_err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "faild to get token",
			})
		}

		uid := sess.Get("uid")
		defer sess.Save()

		if uid == nil {
			return c.Redirect("/auth")
		}

		LS := services.LinkService{}

		links, err := LS.GetLinksOfUser(uid.(string))

		if err != nil {
			fmt.Println(err.Error())
			return c.Render("pages/app", fiber.Map{
				"Title": "Shink - Free Link shortener",
				"UID":   uid,
			})
		}

		return c.Render("pages/app", fiber.Map{
			"Title": "Shink - Free Link shortener",
			"UID":   uid,
			"Links": links,
			"URL":   "sh.mmdjamali.ir",
		})
	})

	app.Get("/auth", func(c *fiber.Ctx) error {
		return c.Render("pages/auth", fiber.Map{
			"Title": "Shink - Authentication",
		})
	})

	app.Get("/confirm", func(c *fiber.Ctx) error {
		return c.Render("pages/confirm-email", fiber.Map{
			"Title": "Shink - Authentication",
			"Email": c.Query("email"),
		})
	})

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/login", controllers.LoginController)

	v1.Post("/confirm", controllers.ValidateOtp)

	links := v1.Group("/links")

	links.Post("/", controllers.CreateLink)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/home")
	})

	app.Get("/:custom", controllers.RedirectController)

	app.Use("/", func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":3001")
}
