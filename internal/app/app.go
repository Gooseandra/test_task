package app

import (
	"github.com/gofiber/fiber/v2"
	"test_task/internal/servises"
)

type App struct {
	Subscribe *servises.Subscribe
	Fetch     *servises.Fetch
	Fib       *fiber.App
	Binance   *servises.Binance
}
