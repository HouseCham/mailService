package middleware

import (
	"github.com/HouseCham/mailService/internal/model"
	"github.com/HouseCham/mailService/internal/validator"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)
// ValidatorMiddleware is a middleware to validate the request body
func ValidatorMiddleware() fiber.Handler {
	log.Info("Validator middleware started")
	return func(c fiber.Ctx) error {
		var request model.UserData
		// Parse the request body from POST request
		if err := c.Bind().JSON(&request); err != nil {
			log.Error("Error parsing request body")
			return badRequestError(c, "Error parsing request body")
		}
		// Validate the request body
		if err := validator.Validate.Struct(request); err != nil {
			log.Error("Error validating request body")
			return badRequestError(c, "Error validating request body")
		}
		return c.Next()
	}
}
// badRequestError is a helper function to return a bad request error
func badRequestError(c fiber.Ctx, errorMessage string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"HasError":  true,
		"Error":     errorMessage,
		"Data":      nil,
		"StatusCode": fiber.StatusBadRequest,
	})
}