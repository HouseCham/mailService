package routes

import (
	"github.com/HouseCham/mailService/api/handlers"
	"github.com/HouseCham/mailService/api/middleware"
	"github.com/gofiber/fiber/v3"
)

// SetRoutes sets up the routes for the application
func SetRoutes(app *fiber.App) {
	mailRoutes := app.Group("/api/mail")
	mailRoutes.Post("/send", handlers.SendMail, middleware.OriginMiddleware(), middleware.ValidatorMiddleware())
}