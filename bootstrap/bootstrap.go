package bootstrap

import (
	"aolus-software/clean-gofiber/adapters/http/routes"
	"aolus-software/clean-gofiber/config"

	"github.com/gofiber/fiber/v2"
)

func AppBootstrap() {
	viper := config.NewViper()
	config := config.AppConfig(viper)

	app := fiber.New(fiber.Config{
		AppName: config.APP_NAME,
	})

	routes.SetupRoutes(app)

	app.Listen(":" + config.APP_PORT)
}