package routes

import (
	"aolus-software/clean-gofiber/adapters/http/handlers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {

	authHandler := handlers.AuthHandler{}

	app.Post("login", authHandler.Login)
	app.Post("register", authHandler.Register)
	app.Post("email-verification", authHandler.EmailVerification)
	app.Post("resend-verification", authHandler.ResendVerification)
	app.Post("forgot-password", authHandler.ForgotPassword)
	app.Post("reset-password", authHandler.ResetPassword)
}