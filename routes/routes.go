package routes

import (
	"warehouse/routes/account"
	"warehouse/routes/vault"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up the routes for the application (and passes the router down to other packages)
func SetupRoutes(router fiber.Router) {

	router.Route("/account", account.SetupRoutes)
	router.Route("/vault", vault.SetupRoutes)

}
