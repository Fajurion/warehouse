package version

import (
	"warehouse/database"
	"warehouse/database/entities"
	"warehouse/util"

	"github.com/gofiber/fiber/v2"
)

func listVersions(c *fiber.Ctx) error {

	// List versions
	var versions []entities.Version
	if err := database.DBConn.Find(&versions).Error; err != nil {
		return util.FailedRequest(c, "failed", nil)
	}

	return c.JSON(fiber.Map{
		"success":  true,
		"versions": versions,
	})
}
