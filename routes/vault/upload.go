package vault

import (
	"log"
	"os"
	"warehouse/util"

	"github.com/gofiber/fiber/v2"
)

// vaultUpload handles the upload of a file to the vault
func vaultUpload(c *fiber.Ctx) error {

	// Parse the request body
	file, err := c.FormFile("file")
	log.Println(file.Filename)
	if err != nil {
		return util.FailedRequest(c, "failed.upload", err)
	}

	err = c.SaveFile(file, os.Getenv("VAULT_PATH")+"/"+file.Filename)
	if err != nil {
		log.Println(err)
		return util.FailedRequest(c, "failed.upload", err)
	}

	return util.SuccessfulRequest(c)
}
