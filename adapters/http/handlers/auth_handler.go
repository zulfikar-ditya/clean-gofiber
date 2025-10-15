package handlers

import "github.com/gofiber/fiber/v2"

type AuthHandler struct {

}

func (a AuthHandler) Login(c *fiber.Ctx) error {
	return c.SendString("Login endpoint")
}

func (a AuthHandler) Register(c *fiber.Ctx) error {
	return c.SendString("Register endpoint")
}

func (a AuthHandler) EmailVerification(c *fiber.Ctx) error {
	return c.SendString("Email verification endpoint")
}

func (a AuthHandler) ResendVerification(c *fiber.Ctx) error {
	return c.SendString("Resend verification endpoint")
}

func (a AuthHandler) ForgotPassword(c *fiber.Ctx) error {
	return c.SendString("Forgot password endpoint")	
}

func (a AuthHandler) ResetPassword(c *fiber.Ctx) error {
	return c.SendString("Reset password endpoint")
}