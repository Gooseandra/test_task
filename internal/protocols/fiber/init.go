package fiber

import (
	"github.com/gofiber/fiber/v2"
	"test_task/internal/app"
)

func Init(app app.App) *fiber.App {
	result := fiber.New()

	result.Post("/subscribe", func(c *fiber.Ctx) error {
		serverAnswer, err := subscribe(c, app)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.SendString(serverAnswer)
	})

	result.Get("/fetch", func(c *fiber.Ctx) error {
		serverAnswer, err := fetchCoins(c, app)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.SendString(serverAnswer)
	})

	return result
}
