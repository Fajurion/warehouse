package auth

import (
	"warehouse/database"
	"warehouse/database/entities"
	"warehouse/util"

	"github.com/gofiber/fiber/v2"
)

type loginRequest struct {
	User     string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func login(c *fiber.Ctx) error {

	// Parse request
	var request loginRequest
	if err := c.BodyParser(&request); err != nil {
		return util.InvalidRequest(c)
	}

	// Validate request
	if err := util.Validator.Struct(request); err != nil {
		return util.InvalidRequest(c)
	}

	// Get account
	var acc entities.Account
	if database.DBConn.Where("username = ?", request.User).Take(&acc).Error != nil {
		return util.FailedRequest(c, "failed.login", nil)
	}

	// Check password
	if acc.Password != util.HashPassword(request.Password) {
		return util.FailedRequest(c, "failed.login", nil)
	}

	// Get role
	var role entities.Role
	if err := database.DBConn.Where("id = ?", acc.Role).Take(&role).Error; err != nil {
		return util.FailedRequest(c, "server.error", err)
	}

	// Create token
	token, err := util.Token(acc.Role, util.TokenExpiration(), acc.ID)
	if err != nil {
		return util.FailedRequest(c, "server.error", err)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}
