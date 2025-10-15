package routes

import (
	"aolus-software/clean-gofiber/adapters/http/handlers"

	"github.com/gofiber/fiber/v2"
)

func ProfileRoutes(app *fiber.App) {
	profileHandler := handlers.ProfileHandler{}

	app.Get("/profile", profileHandler.GetProfile)
	app.Post("/logout", profileHandler.Logout)
	app.Post("/refresh-token", profileHandler.RefreshToken)
}