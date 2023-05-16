package version

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router) {
	router.Post("/create", createVersion)
	router.Post("/publish", publishVersion)
	router.Get("/list", listVersions)
}
