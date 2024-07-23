package handlers

import "github.com/gofiber/fiber/v3"

// Handler function for sending mail
func SendMail(c fiber.Ctx) error {
	// Send mail logic here
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Mail sent successfully",
	})
}