package branch

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router) {
	router.Post("/create", createBranch)
	router.Get("/list", listBranches)
}
