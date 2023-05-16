package version

import (
	"warehouse/database"
	"warehouse/database/entities"
	"warehouse/util"

	"github.com/gofiber/fiber/v2"
)

type createRequest struct {
	Name   string `json:"name" validate:"required,max=32"`
	Branch uint   `json:"branch" validate:"required"`
}

func createVersion(c *fiber.Ctx) error {

	// Parse request
	var req createRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	// Validate request
	if err := util.Validator.Struct(req); err != nil {
		return util.InvalidRequest(c)
	}

	// Get branch
	var branch entities.Branch
	if database.DBConn.Where("id = ?", req.Branch).First(&branch).Error != nil {
		return util.FailedRequest(c, "not.found", nil)
	}

	// Create version
	var version entities.Version = entities.Version{
		ID:     util.GenerateToken(24),
		Name:   req.Name,
		Branch: branch.ID,
	}

	if database.DBConn.Create(&version).Error != nil {
		return util.FailedRequest(c, "failed", nil)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"version": version,
	})
}
