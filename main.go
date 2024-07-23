package main

import (
	"fmt"
	"os"

	"github.com/HouseCham/mailService/api/routes"
	"github.com/HouseCham/mailService/internal/validator"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	log.Info("Starting mail service")
	// Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }
	// Creating a new Fiber app
	log.Info("Setting up Fiber instance with CORS middleware and routes")
	app := fiber.New()
	// Setting up CORS middleware
	app.Use((cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("ALLOWED_ORIGIN")},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Origin, Content-Type"},
	})))
	log.Info("Setting up routes")
	// Setting up routes
	routes.SetRoutes(app)
	log.Info("Setting up validator")
	// Setting up validator
	validator.SetUpValidations()
	// Starting the server
	log.Infof("Server is running on http://%s:%d", os.Getenv("HOST"), os.Getenv("PORT"))
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}