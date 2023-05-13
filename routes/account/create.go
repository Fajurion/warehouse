package account

import (
	"os"
	"warehouse/database"
	"warehouse/database/entities"
	"warehouse/util"

	"github.com/gofiber/fiber/v2"
)

type createRequest struct {

	// Admin token
	AdminToken string `json:"admin_token" validate:"required"`

	// Details of the account
	Username string `json:"username" validate:"required,max=32,min=3"`
}

// createAccount creates an account for the panel
func createAccount(c *fiber.Ctx) error {

	// Parse request
	var request createRequest
	if err := c.BodyParser(&request); err != nil {
		return util.InvalidRequest(c)
	}

	// Validate request
	err := util.Validator.Struct(request)
	if err != nil {
		return util.InvalidRequest(c)
	}

	// Check if admin token is valid
	if request.AdminToken != os.Getenv("ADMIN_TOKEN") {
		return util.InvalidRequest(c)
	}

	// Create account
	pw := util.GenerateToken(32)
	database.DBConn.Create(&entities.Account{
		Username: request.Username,
		Password: util.HashPassword(pw),
	})

	// Return password
	return c.Status(200).JSON(fiber.Map{
		"success":  true,
		"password": pw,
	})
}
