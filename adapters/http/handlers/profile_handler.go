package handlers

import "github.com/gofiber/fiber/v2"	

type ProfileHandler struct{

}

func (p ProfileHandler) GetProfile(c *fiber.Ctx) error {
	return c.SendString("Get profile endpoint")
}

func (p ProfileHandler) UpdateProfile(c *fiber.Ctx) error {
	return c.SendString("Update profile endpoint")
}

func (p ProfileHandler) UpdatePassword(c *fiber.Ctx) error {
	return c.SendString("Update password endpoint")
}

func (a ProfileHandler) Logout(c *fiber.Ctx) error {
	return c.SendString("Logout endpoint")
}

func (a ProfileHandler) RefreshToken(c *fiber.Ctx) error {
	return c.SendString("Refresh token endpoint")
}