package v1

import (
	"git.recombee.net/jtrojak/server-db/constants"
	"github.com/gofiber/fiber/v2"
)

type Base struct {
}

func (b *Base) ping(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"message": "pong",
	})
}

func (b *Base) app(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"app":     constants.AppName,
		"version": constants.Version,
	})
}
