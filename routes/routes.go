package routes

import (
	"os"
	"warehouse/routes/account"
	"warehouse/routes/app"
	"warehouse/routes/auth"
	"warehouse/routes/vault"
	"warehouse/util"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// SetupRoutes sets up the routes for the application (and passes the router down to other packages)
func SetupRoutes(router fiber.Router) {

	router.Route("/account", account.SetupRoutes)
	router.Route("/auth", auth.SetupRoutes)

	// Authorized routes
	router.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),

		// Checks if the token is expired
		SuccessHandler: func(c *fiber.Ctx) error {

			if util.IsExpired(c) {
				return util.InvalidRequest(c)
			}

			return c.Next()
		},

		// Error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			// Return error message
			return c.SendStatus(401)
		},
	}))

	router.Route("/vault", vault.SetupRoutes)
	router.Route("/app", app.SetupRoutes)

}
