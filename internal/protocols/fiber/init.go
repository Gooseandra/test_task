package fiber

import (
	"github.com/gofiber/fiber/v2"
	"test_task/internal/app"
)

func Init(app app.App) *fiber.App {
	result := fiber.New()

	result.Post("/subscribe", func(c *fiber.Ctx) error {
		return c.SendString(subscribe(c, app))
	})

	result.Get("/fetch", func(c *fiber.Ctx) error {
		return c.SendString(fetch(c, app))
	})

	return result
}
