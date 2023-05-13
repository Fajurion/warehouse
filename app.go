package main

import (
	"os"
	"warehouse/database"
	"warehouse/database/credentials"
	"warehouse/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Load environment variables
	credentials.Load()

	// Connect to the database
	database.Connect()

	testMode := os.Getenv("TEST_MODE") == "true"

	// Fiber
	app := fiber.New()

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
