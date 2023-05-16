package branch

import (
	"warehouse/database"
	"warehouse/database/entities"
	"warehouse/util"

	"github.com/gofiber/fiber/v2"
)

func listBranches(c *fiber.Ctx) error {

	// List branches
	var branches []entities.Branch
	if err := database.DBConn.Find(&branches).Error; err != nil {
		return util.FailedRequest(c, "failed.list", nil)
	}

	return c.JSON(fiber.Map{
		"success":  true,
		"branches": branches,
	})
}
