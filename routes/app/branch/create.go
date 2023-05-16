package branch

import (
	"warehouse/database"
	"warehouse/database/entities"
	"warehouse/util"

	"github.com/gofiber/fiber/v2"
)

type createRequest struct {
	Name        string `json:"name" validate:"required,max=32"`
	App         uint   `json:"app" validate:"required"`
	Description string `json:"description" validate:"required,max=255"`
}

func createBranch(c *fiber.Ctx) error {

	// Parse request
	var req createRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	// Validate request
	if err := util.Validator.Struct(req); err != nil {
		return util.InvalidRequest(c)
	}

	// Get app
	var app entities.App
	if database.DBConn.Where("id = ?", req.App).First(&app).Error != nil {
		return util.FailedRequest(c, "failed.get", nil)
	}

	// Create branch
	var branch entities.Branch = entities.Branch{
		Name:        req.Name,
		App:         app.ID,
		Description: req.Description,
	}

	if database.DBConn.Create(&branch).Error != nil {
		return util.FailedRequest(c, "failed.create", nil)
	}

	return util.SuccessfulRequest(c)
}
