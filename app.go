package main

import (
	"os"
	"warehouse/database"
	"warehouse/database/credentials"
	"warehouse/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	// Load environment variables
	credentials.Load()

	// Connect to the database
	database.Connect()

	testMode := os.Getenv("TEST_MODE") == "true"

	// Fiber
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	app.Route("/", routes.SetupRoutes)

	// Listen on default port
	port := os.Getenv("PORT")
	if testMode {
		app.Listen("localhost:" + port)
	} else {
		app.Listen("0.0.0.0:" + port)
	}

}
