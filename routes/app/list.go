package app

import (
	"warehouse/database"
	"warehouse/database/entities"
	"warehouse/util"

	"github.com/gofiber/fiber/v2"
)

func listApps(c *fiber.Ctx) error {

	// List apps
	var apps []entities.App
	if err := database.DBConn.Find(&apps).Error; err != nil {
		return util.FailedRequest(c, "failed.list", nil)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"apps":    apps,
	})
}
