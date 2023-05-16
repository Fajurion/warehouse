package app

import (
	"warehouse/database"
	"warehouse/database/entities"
	"warehouse/util"

	"github.com/gofiber/fiber/v2"
)

type createRequest struct {
	Name        string `json:"name" validate:"required,max=32"`
	Description string `json:"description" validate:"required,max=255"`
}

func createApp(c *fiber.Ctx) error {

	// Parse request
	var req createRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	// Validate request
	if err := util.Validator.Struct(req); err != nil {
		return util.InvalidRequest(c)
	}

	// Create app
	var app entities.App = entities.App{
		Name:        req.Name,
		Description: req.Description,
	}

	if database.DBConn.Create(&app).Error != nil {
		return util.FailedRequest(c, "failed.create", nil)
	}

	return util.SuccessfulRequest(c)
}
