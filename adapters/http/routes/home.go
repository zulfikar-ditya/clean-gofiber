package routes

import (
	"aolus-software/clean-gofiber/internal/config"
	"time"

	"github.com/gofiber/fiber/v2"
)

func HomeRoutes(app *fiber.App) {
	startTime := time.Now()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"app_name": config.Env("APP_NAME"),
			"timestamp": time.Now().Unix(),
		})
	})

	app.Get("/uptime", func(c *fiber.Ctx) error {
		uptime := time.Since(startTime)
		return c.JSON(fiber.Map{
			"uptime": uptime.String(),
		})
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"timestamp": time.Now().Unix(),
		})
	})
}