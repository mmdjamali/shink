package sessions

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var SessionStore *session.Store

func InitilaizeSessionStore() {
	SessionStore = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     30 * 24 * time.Hour,
	})

	SessionStore.RegisterType(fiber.Map{})
}
