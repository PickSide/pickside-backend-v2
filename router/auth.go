package router

import (
	"github.com/gofiber/fiber/v2"
)

func Connect(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func Disconnect(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func RefreshToken(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func Register(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
