package app

import (
	"warehouse/routes/app/branch"
	"warehouse/routes/app/version"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	router.Route("/version", version.SetupRoutes)
	router.Route("/branch", branch.SetupRoutes)

	router.Post("/create", createApp)
	router.Get("/list", listApps)
}
