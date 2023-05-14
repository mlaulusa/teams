package middleware

import (
	"teams/env"
	"teams/session"

	"github.com/gofiber/fiber/v2"
)

func AddGuard(app *fiber.App) {
}

func authenticate(c *fiber.Ctx) error {
	store := c.Locals(env.SESSION).(session.Store)
	sesh, err := store.Get(c)

	if err != nil {
		return err
	}

	accountId := sesh.Get("account_id")

	if accountId == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
