package version

import (
	"warehouse/database"
	"warehouse/database/entities"
	"warehouse/util"

	"github.com/gofiber/fiber/v2"
)

type publishRequest struct {
	Version string `json:"version" validate:"required"`
}

func publishVersion(c *fiber.Ctx) error {

	// Parse request
	var req publishRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	// Validate request
	if err := util.Validator.Struct(req); err != nil {
		return util.InvalidRequest(c)
	}

	// Publish version
	var version entities.Version
	if database.DBConn.Where("id = ?", req.Version).First(&version).Error != nil {
		return util.FailedRequest(c, "not.found", nil)
	}

	// Update branch
	if database.DBConn.Model(&entities.Branch{}).Where("id = ?", version.Branch).
		Update("latest_version", version.ID).Error != nil {

		return util.FailedRequest(c, "failed", nil)
	}

	return util.SuccessfulRequest(c)
}
