package main

import (
	"bufio"
	"log"
	"os"
	"warehouse/database"
	"warehouse/database/credentials"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Load environment variables
	credentials.Load()

	// Connect to the database
	database.Connect()

	scanner := bufio.NewScanner(os.Stdin)

	log.Println("Do you want to run this node in testing mode? (y/n)")

	scanner.Scan()
	testMode := scanner.Text() == "y"

	// Fiber
	app := fiber.New()

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Listen on default port
	port := os.Getenv("PORT")
	if testMode {
		app.Listen("localhost:" + port)
	} else {
		app.Listen("0.0.0.0:" + port)
	}

}
